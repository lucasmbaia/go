package pubnub

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pubnub/go/utils"
)

const DELETE_CHANNEL_GROUP = "/v1/channel-registration/sub-key/%s/channel-group/%s/remove"

var emptyDeleteChannelGroupResponse *DeleteChannelGroupResponse

type deleteChannelGroupBuilder struct {
	opts *deleteChannelGroupOpts
}

func newDeleteChannelGroupBuilder(pubnub *PubNub) *deleteChannelGroupBuilder {
	builder := deleteChannelGroupBuilder{
		opts: &deleteChannelGroupOpts{
			pubnub: pubnub,
		},
	}

	return &builder
}

func newDeleteChannelGroupBuilderWithContext(
	pubnub *PubNub, context Context) *deleteChannelGroupBuilder {
	builder := deleteChannelGroupBuilder{
		opts: &deleteChannelGroupOpts{
			pubnub: pubnub,
			ctx:    context,
		},
	}

	return &builder
}

func (b *deleteChannelGroupBuilder) ChannelGroup(
	cg string) *deleteChannelGroupBuilder {
	b.opts.ChannelGroup = cg
	return b
}

func (b *deleteChannelGroupBuilder) Execute() (
	*DeleteChannelGroupResponse, StatusResponse, error) {
	_, status, err := executeRequest(b.opts)
	if err != nil {
		return emptyDeleteChannelGroupResponse, status, err
	}

	return emptyDeleteChannelGroupResponse, status, nil
}

type deleteChannelGroupOpts struct {
	pubnub *PubNub

	ChannelGroup string

	Transport http.RoundTripper

	ctx Context
}

func (o *deleteChannelGroupOpts) config() Config {
	return *o.pubnub.Config
}

func (o *deleteChannelGroupOpts) client() *http.Client {
	return o.pubnub.GetClient()
}

func (o *deleteChannelGroupOpts) context() Context {
	return o.ctx
}

func (o *deleteChannelGroupOpts) validate() error {
	if o.config().SubscribeKey == "" {
		return newValidationError(o, StrMissingSubKey)
	}

	if o.ChannelGroup == "" {
		return newValidationError(o, StrMissingChannelGroup)
	}

	return nil
}

type DeleteChannelGroupResponse struct{}

func (o *deleteChannelGroupOpts) buildPath() (string, error) {
	return fmt.Sprintf(DELETE_CHANNEL_GROUP,
		o.pubnub.Config.SubscribeKey,
		utils.UrlEncode(o.ChannelGroup)), nil
}

func (o *deleteChannelGroupOpts) buildQuery() (*url.Values, error) {
	q := defaultQuery(o.pubnub.Config.Uuid)

	return q, nil
}

func (o *deleteChannelGroupOpts) buildBody() ([]byte, error) {
	return []byte{}, nil
}

func (o *deleteChannelGroupOpts) httpMethod() string {
	return "GET"
}

func (o *deleteChannelGroupOpts) isAuthRequired() bool {
	return true
}

func (o *deleteChannelGroupOpts) requestTimeout() int {
	return o.pubnub.Config.NonSubscribeRequestTimeout
}

func (o *deleteChannelGroupOpts) connectTimeout() int {
	return o.pubnub.Config.ConnectTimeout
}

func (o *deleteChannelGroupOpts) operationType() OperationType {
	return PNRemoveGroupOperation
}