package upbit

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	UpbitKR = "https://api.upbit.com"
	UpbitSG = "https://sg-api.upbit.com"
	UpbitID = "https://id-api.upbit.com"
)

type ClientOptions struct {
	AccessKey string
	SecretKey string
	ServerURL string
}

func ClientOptionsFromEnv() *ClientOptions {
	return &ClientOptions{
		SecretKey: os.Getenv("UPBIT_OPEN_API_SECRET_KEY"),
		ServerURL: os.Getenv("UPBIT_OPEN_API_SERVER_URL"),
		AccessKey: os.Getenv("UPBIT_OPEN_API_ACCESS_KEY"),
	}
}

type service struct {
	client *Client
}

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	common     service

	accessKey string
	secretKey string

	Accounts  *AccountService
	Orders    *OrderService
	Withdraws *WithdrawService
	Deposits  *DepositService
	Markets   *MarketService
	Candles   *CandleService
}

type (
	AccountService  service
	MarketService   service
	OrderService    service
	WithdrawService service
	DepositService  service
	CandleService   service
)

func NewClient(httpClient *http.Client, opt *ClientOptions) (*Client, error) {
	serverURL := UpbitKR
	if opt.ServerURL != "" {
		serverURL = opt.ServerURL
	}

	baseURL, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		accessKey:  opt.AccessKey,
		secretKey:  opt.SecretKey,
		baseURL:    baseURL,
		httpClient: httpClient,
	}

	c.common.client = c
	c.Accounts = (*AccountService)(&c.common)
	c.Orders = (*OrderService)(&c.common)
	c.Withdraws = (*WithdrawService)(&c.common)
	c.Deposits = (*DepositService)(&c.common)
	c.Markets = (*MarketService)(&c.common)
	c.Candles = (*CandleService)(&c.common)
	return c, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	if ctx == nil {
		ctx = context.TODO()
	}

	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

var (
	ErrNotImplemented   = fmt.Errorf("upbit: not implemented")
	ErrInvalidArguments = fmt.Errorf("upbit: invalid arguments")
)

// API document doesn't specifiy error model. This might change.
type ErrResponse struct {
	Detail struct {
		Message string `json:"message"`
		Name    string `json:"name"`
	} `json:"error"`
}

func (e *ErrResponse) Error() string {
	return fmt.Sprintf("upbit: %s, %s", e.Detail.Name, e.Detail.Message)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	errResp := &ErrResponse{}
	err = json.Unmarshal(body, errResp)

	// If server response with sturctured error
	if err == nil && errResp.Detail.Name != "" {
		return errResp
	}

	return errors.New(string(body))
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func (c *Client) generateToken(req *http.Request, queryString string) error {
	queryString, err := url.QueryUnescape(queryString)
	if err != nil {
		return err
	}

	h := sha512.New()
	_, err = h.Write([]byte(queryString))
	if err != nil {
		return err
	}

	qhs := hex.EncodeToString(h.Sum(nil))

	nonce, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	var tokenString string

	if queryString != "" {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"access_key":     c.accessKey,
			"nonce":          nonce.String(),
			"query_hash":     qhs,
			"query_hash_alg": "SHA512",
		})

		tokenString, err = token.SignedString([]byte(c.secretKey))
		if err != nil {
			return nil
		}
	} else {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"access_key": c.accessKey,
			"nonce":      nonce.String(),
		})

		tokenString, err = token.SignedString([]byte(c.secretKey))
		if err != nil {
			return nil
		}
	}

	req.Header.Add("Authorization", "Bearer "+tokenString)
	return nil
}
