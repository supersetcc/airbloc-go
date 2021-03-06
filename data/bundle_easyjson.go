// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package data

import (
	json "encoding/json"

	common "github.com/airbloc/airbloc-go/common"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson40af25b0DecodeGithubComAirblocAirblocGoData(in *jlexer.Lexer, out *Bundle) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "provider":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Provider).UnmarshalJSON(data))
			}
		case "collection":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Collection).UnmarshalJSON(data))
			}
		case "dataCount":
			out.DataCount = int(in.Int())
		case "ingestedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.IngestedAt).UnmarshalJSON(data))
			}
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]*common.EncryptedData, 0, 8)
					} else {
						out.Data = []*common.EncryptedData{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *common.EncryptedData
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(common.EncryptedData)
						}
						easyjson40af25b0DecodeGithubComAirblocAirblocGoCommon(in, &*v1)
					}
					out.Data = append(out.Data, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson40af25b0EncodeGithubComAirblocAirblocGoData(out *jwriter.Writer, in Bundle) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"provider\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Provider).MarshalJSON())
	}
	{
		const prefix string = ",\"collection\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Collection).MarshalJSON())
	}
	{
		const prefix string = ",\"dataCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.DataCount))
	}
	{
		const prefix string = ",\"ingestedAt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.IngestedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Data {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson40af25b0EncodeGithubComAirblocAirblocGoCommon(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bundle) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson40af25b0EncodeGithubComAirblocAirblocGoData(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bundle) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson40af25b0EncodeGithubComAirblocAirblocGoData(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bundle) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson40af25b0DecodeGithubComAirblocAirblocGoData(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bundle) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson40af25b0DecodeGithubComAirblocAirblocGoData(l, v)
}
func easyjson40af25b0DecodeGithubComAirblocAirblocGoCommon(in *jlexer.Lexer, out *common.EncryptedData) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ownerAnid":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.OwnerAnid).UnmarshalJSON(data))
			}
		case "payload":
			if in.IsNull() {
				in.Skip()
				out.Payload = nil
			} else {
				out.Payload = in.Bytes()
			}
		case "capsule":
			if in.IsNull() {
				in.Skip()
				out.Capsule = nil
			} else {
				out.Capsule = in.Bytes()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson40af25b0EncodeGithubComAirblocAirblocGoCommon(out *jwriter.Writer, in common.EncryptedData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ownerAnid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.OwnerAnid).MarshalJSON())
	}
	{
		const prefix string = ",\"payload\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Payload)
	}
	{
		const prefix string = ",\"capsule\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Capsule)
	}
	out.RawByte('}')
}
