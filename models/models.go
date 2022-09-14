package models

import "encoding/json"

func UnmarshalMessageResponse(data []byte) (MessageResponse, error) {
	var r MessageResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type MessageResponse struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string    `json:"messaging_product"`
	Metadata         Metadata  `json:"metadata"`
	Contacts         []Contact `json:"contacts"`
	Messages         []Message `json:"messages"`
}

type Contact struct {
	Profile Profile `json:"profile"`
	WaID    string  `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Message struct {
	From      string `json:"from"`
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Image     Image  `json:"image"`
}

type Image struct {
	MIMEType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

func UnmarshalMediaResponse(data []byte) (MediaResponse, error) {
	var r MediaResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type MediaResponse struct {
	URL              string `json:"url"`
	MIMEType         string `json:"mime_type"`
	Sha256           string `json:"sha256"`
	FileSize         int64  `json:"file_size"`
	ID               string `json:"id"`
	MessagingProduct string `json:"messaging_product"`
}

func (r *StickerRequest) MarshalStickerRequest() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStickerRequest(data []byte) (StickerRequest, error) {
	var r StickerRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

type StickerRequest struct {
	MessagingProduct string  `json:"messaging_product"`
	RecipientType    string  `json:"recipient_type"`
	To               string  `json:"to"`
	Type             string  `json:"type"`
	Sticker          Sticker `json:"sticker"`
}

type Sticker struct {
	ID string `json:"id"`
}

func UnmarshalTextMessageRequest(data []byte) (TextMessageRequest, error) {
	var r TextMessageRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TextMessageRequest) MarshalTextMessageRequest() ([]byte, error) {
	return json.Marshal(r)
}

type TextMessageRequest struct {
	MessagingProduct string `json:"messaging_product"`
	RecipientType    string `json:"recipient_type"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             Text   `json:"text"`
}

type Text struct {
	PreviewURL bool   `json:"preview_url"`
	Body       string `json:"body"`
}
