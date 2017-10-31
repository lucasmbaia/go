package pubnub

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pubnub/go/utils"
)

const ADD_CHANNEL_CHANNEL_GROUP_PATH = "/v1/channel-registration/sub-key/%s/channel-group/%s"

var emptyAddChannelToChannelGroupResp *AddChannelToChannelGroupResponse

type AddChannelToChannelGroupBuilder struct {
	opts *addChannelOpts
}

func newAddChannelToChannelGroupBuilder(
	pubnub *PubNub) *AddChannelToChannelGroupBuilder {
	builder := AddChannelToChannelGroupBuilder{
		opts: &addChannelOpts{
			pubnub: pubnub,
		},
	}

	return &builder
}

func newAddChannelToChannelGroupBuilderWithContext(
	pubnub *PubNub, context Context) *AddChannelToChannelGroupBuilder {
	builder := AddChannelToChannelGroupBuilder{
		opts: &addChannelOpts{
			pubnub: pubnub,
			ctx:    context,
		},
	}

	return &builder
}

func (b *AddChannelToChannelGroupBuilder) Channels(
	ch []string) *AddChannelToChannelGroupBuilder {
	b.opts.Channels = ch

	return b
}

func (b *AddChannelToChannelGroupBuilder) Group(
	cg string) *AddChannelToChannelGroupBuilder {
	b.opts.Group = cg

	return b
}

func (b *AddChannelToChannelGroupBuilder) Transport(
	tr http.RoundTripper) *AddChannelToChannelGroupBuilder {
	b.opts.Transport = tr

	return b
}

func (b *AddChannelToChannelGroupBuilder) Execute() (
	*AddChannelToChannelGroupResponse, StatusResponse, error) {
	rawJson, status, err := executeRequest(b.opts)
	if err != nil {
		return emptyAddChannelToChannelGroupResp, status, err
	}

	return newAddChannelToChannelGroupsResponse(rawJson, status)
}

type addChannelOpts struct {
	pubnub *PubNub

	Channels []string

	Group string

	Transport http.RoundTripper

	ctx Context
}

func (o *addChannelOpts) config() Config {
	return *o.pubnub.Config
}

func (o *addChannelOpts) client() *http.Client {
	return o.pubnub.GetClient()
}

func (o *addChannelOpts) context() Context {
	return o.ctx
}

func (o *addChannelOpts) validate() error {
	if o.config().SubscribeKey == "" {
		return newValidationError(o, StrMissingSubKey)
	}

	if len(o.Channels) == 0 {
		return newValidationError(o, StrMissingChannel)
	}

	if o.Group == "" {
		return newValidationError(o, StrMissingChannelGroup)
	}

	return nil
}

func (o *addChannelOpts) buildPath() (string, error) {
	return fmt.Sprintf(ADD_CHANNEL_CHANNEL_GROUP_PATH,
		o.pubnub.Config.SubscribeKey,
		utils.UrlEncode(o.Group)), nil
}

func (o *addChannelOpts) buildQuery() (*url.Values, error) {
	q := defaultQuery(o.pubnub.Config.Uuid)

	q.Set("add", strings.Join(o.Channels, ","))

	return q, nil
}

func (o *addChannelOpts) buildBody() ([]byte, error) {
	return []byte{}, nil
}

func (o *addChannelOpts) httpMethod() string {
	return "GET"
}

func (o *addChannelOpts) isAuthRequired() bool {
	return true
}

func (o *addChannelOpts) requestTimeout() int {
	return o.pubnub.Config.NonSubscribeRequestTimeout
}

func (o *addChannelOpts) connectTimeout() int {
	return o.pubnub.Config.ConnectTimeout
}

func (o *addChannelOpts) operationType() OperationType {
	return PNAddChannelsToChannelGroupOperation
}

type AddChannelToChannelGroupResponse struct {
}

func newAddChannelToChannelGroupsResponse(jsonBytes []byte, status StatusResponse) (
	*AddChannelToChannelGroupResponse, StatusResponse, error) {

	return emptyAddChannelToChannelGroupResp, status, nil
}
