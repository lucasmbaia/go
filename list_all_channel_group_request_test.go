package pubnub

import (
	"fmt"
	"net/url"
	"testing"

	h "github.com/pubnub/go/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func TestListAllChannelGroupRequestBasic(t *testing.T) {
	assert := assert.New(t)

	opts := &allChannelGroupOpts{
		ChannelGroup: "cg",
		pubnub:       pubnub,
	}

	path, err := opts.buildPath()
	assert.Nil(err)
	u := &url.URL{
		Path: path,
	}
	h.AssertPathsEqual(t,
		fmt.Sprintf("/v1/channel-registration/sub-key/sub_key/channel-group/cg"),
		u.EscapedPath(), []int{})

	query, err := opts.buildQuery()
	assert.Nil(err)

	expected := &url.Values{}

	h.AssertQueriesEqual(t, expected, query, []string{"pnsdk", "uuid"}, []string{})

	body, err := opts.buildBody()
	assert.Nil(err)
	assert.Equal([]byte{}, body)
}

func TestListAllChannelsNewAllChannelGroupResponseErrorUnmarshalling(t *testing.T) {
	assert := assert.New(t)
	jsonBytes := []byte(`s`)

	_, _, err := newAllChannelGroupResponse(jsonBytes, StatusResponse{})
	assert.Equal("pubnub/parsing: Error unmarshalling response: {s}", err.Error())
}

func TestListAllChannelsValidateSubscribeKey(t *testing.T) {
	assert := assert.New(t)
	pn := NewPubNub(NewDemoConfig())
	pn.Config.SubscribeKey = ""
	opts := &allChannelGroupOpts{
		ChannelGroup: "cg",
		pubnub:       pn,
	}

	assert.Equal("pubnub/validation: pubnub: \x0f: Missing Subscribe Key", opts.validate().Error())
}

func TestListAllChannelsValidateChannelGrp(t *testing.T) {
	assert := assert.New(t)
	pn := NewPubNub(NewDemoConfig())
	opts := &allChannelGroupOpts{
		pubnub: pn,
	}

	assert.Equal("pubnub/validation: pubnub: \x0f: Missing Channel Group", opts.validate().Error())
}
