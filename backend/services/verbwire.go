package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// VerbwireService handles all Verbwire API interactions
type VerbwireService struct {
	APIKey    string
	PublicKey string
	BaseURL   string
	Chain     string
}

// NewVerbwireService creates a new Verbwire service instance
func NewVerbwireService() *VerbwireService {
	return &VerbwireService{
		APIKey:    os.Getenv("VERBWIRE_API_KEY"),
		PublicKey: os.Getenv("VERBWIRE_PUBLIC_KEY"),
		BaseURL:   os.Getenv("VERBWIRE_BASE_URL"),
		Chain:     os.Getenv("CHAIN"),
	}
}

// MintNFTRequest represents the request to mint an NFT
type MintNFTRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ImageURL        string `json:"imageUrl"`
	RecipientAddress string `json:"recipientAddress"`
	Chain           string `json:"chain"`
	Quantity        int    `json:"quantity"`
}

// MintNFTResponse represents the response from minting an NFT
type MintNFTResponse struct {
	Success         bool   `json:"success"`
	TransactionHash string `json:"transaction_hash"`
	ContractAddress string `json:"contract_address"`
	TokenID         string `json:"token_id"`
	OpenseaURL      string `json:"opensea_url"`
	Message         string `json:"message"`
	Details         struct {
		TransactionHash  string `json:"transactionHash"`
		TransactionIndex int    `json:"transactionIndex"`
		BlockHash        string `json:"blockHash"`
		BlockNumber      int    `json:"blockNumber"`
		From             string `json:"from"`
		To               string `json:"to"`
		ContractAddress  string `json:"contractAddress"`
		TokenID          string `json:"tokenId"`
	} `json:"details"`
}

// QuickMintNFT mints an NFT using Verbwire's Quick Mint API
func (v *VerbwireService) QuickMintNFT(req MintNFTRequest) (*MintNFTResponse, error) {
	url := fmt.Sprintf("%s/nft/mint/quickMintFromMetadata", v.BaseURL)
	
	// Create multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	
	// Create metadata object
	metadata := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"image":       req.ImageURL,
	}
	metadataJSON, _ := json.Marshal(metadata)
	
	// Add form fields
	writer.WriteField("chain", v.Chain)
	writer.WriteField("data", string(metadataJSON))
	writer.WriteField("recipientAddress", req.RecipientAddress)
	
	err := writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}
	
	// Create HTTP request
	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	// Set headers
	httpReq.Header.Set("X-API-Key", v.APIKey)
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Accept", "application/json")
	
	// Execute request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	
	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	
	// Parse response
	var mintResp MintNFTResponse
	if err := json.Unmarshal(respBody, &mintResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, mintResp.Message)
	}
	
	return &mintResp, nil
}

// GetNFTsByWallet retrieves all NFTs owned by a wallet address
func (v *VerbwireService) GetNFTsByWallet(walletAddress string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/data/nftsByWalletAddress", v.BaseURL)
	
	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	// Add query parameters
	q := req.URL.Query()
	q.Add("walletAddress", walletAddress)
	q.Add("chain", v.Chain)
	req.URL.RawQuery = q.Encode()
	
	// Set headers
	req.Header.Set("X-API-Key", v.APIKey)
	req.Header.Set("Accept", "application/json")
	
	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	
	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	
	// Parse response
	var result struct {
		NFTs []map[string]interface{} `json:"nfts"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	
	return result.NFTs, nil
}

// GetNFTMetadata retrieves metadata for a specific NFT
func (v *VerbwireService) GetNFTMetadata(contractAddress, tokenID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/nft/data/nftDetails", v.BaseURL)
	
	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	// Add query parameters
	q := req.URL.Query()
	q.Add("contractAddress", contractAddress)
	q.Add("tokenId", tokenID)
	q.Add("chain", v.Chain)
	req.URL.RawQuery = q.Encode()
	
	// Set headers
	req.Header.Set("X-API-Key", v.APIKey)
	req.Header.Set("Accept", "application/json")
	
	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	
	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	
	// Parse response
	var metadata map[string]interface{}
	if err := json.Unmarshal(respBody, &metadata); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	
	return metadata, nil
}

// TransferNFT transfers an NFT from one wallet to another
func (v *VerbwireService) TransferNFT(contractAddress, tokenID, fromAddress, toAddress string) (*map[string]interface{}, error) {
	url := fmt.Sprintf("%s/nft/transfer", v.BaseURL)
	
	// Create multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	
	// Add form fields
	writer.WriteField("chain", v.Chain)
	writer.WriteField("contractAddress", contractAddress)
	writer.WriteField("tokenId", tokenID)
	writer.WriteField("fromAddress", fromAddress)
	writer.WriteField("toAddress", toAddress)
	
	err := writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}
	
	// Create HTTP request
	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	// Set headers
	httpReq.Header.Set("X-API-Key", v.APIKey)
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Accept", "application/json")
	
	// Execute request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	
	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	
	// Parse response
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	
	return &result, nil
}

// GetCollectionStats retrieves statistics for an NFT collection
func (v *VerbwireService) GetCollectionStats(contractAddress string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/nft/data/collectionStatistics", v.BaseURL)
	
	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	// Add query parameters
	q := req.URL.Query()
	q.Add("contractAddress", contractAddress)
	q.Add("chain", v.Chain)
	req.URL.RawQuery = q.Encode()
	
	// Set headers
	req.Header.Set("X-API-Key", v.APIKey)
	req.Header.Set("Accept", "application/json")
	
	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	
	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	
	// Parse response
	var stats map[string]interface{}
	if err := json.Unmarshal(respBody, &stats); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	
	return stats, nil
}
