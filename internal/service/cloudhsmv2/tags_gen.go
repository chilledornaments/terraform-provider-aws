// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package cloudhsmv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2/cloudhsmv2iface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
)

// ListTags lists cloudhsmv2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn cloudhsmv2iface.CloudHSMV2API, identifier string) (tftags.KeyValueTags, error) {
	input := &cloudhsmv2.ListTagsInput{
		ResourceId: aws.String(identifier),
	}

	output, err := conn.ListTagsWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.TagList), nil
}

func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := ListTags(ctx, meta.(*conns.AWSClient).CloudHSMV2Conn(), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns cloudhsmv2 service tags.
func Tags(tags tftags.KeyValueTags) []*cloudhsmv2.Tag {
	result := make([]*cloudhsmv2.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &cloudhsmv2.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from cloudhsmv2 service tags.
func KeyValueTags(ctx context.Context, tags []*cloudhsmv2.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// GetTagsIn returns cloudhsmv2 service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*cloudhsmv2.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets cloudhsmv2 service tags in Context.
func SetTagsOut(ctx context.Context, tags []*cloudhsmv2.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// UpdateTags updates cloudhsmv2 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.

func UpdateTags(ctx context.Context, conn cloudhsmv2iface.CloudHSMV2API, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &cloudhsmv2.UntagResourceInput{
			ResourceId: aws.String(identifier),
			TagKeyList: aws.StringSlice(removedTags.IgnoreAWS().Keys()),
		}

		_, err := conn.UntagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &cloudhsmv2.TagResourceInput{
			ResourceId: aws.String(identifier),
			TagList:    Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.TagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).CloudHSMV2Conn(), identifier, oldTags, newTags)
}
