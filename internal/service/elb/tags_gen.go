// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package elb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
)

// ListTags lists elb service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn elbiface.ELBAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &elb.DescribeTagsInput{
		LoadBalancerNames: aws.StringSlice([]string{identifier}),
	}

	output, err := conn.DescribeTagsWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.TagDescriptions[0].Tags), nil
}

func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := ListTags(ctx, meta.(*conns.AWSClient).ELBConn(), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// TagKeys returns elb service tag keys.
func TagKeys(tags tftags.KeyValueTags) []*elb.TagKeyOnly {
	result := make([]*elb.TagKeyOnly, 0, len(tags))

	for k := range tags.Map() {
		tagKey := &elb.TagKeyOnly{
			Key: aws.String(k),
		}

		result = append(result, tagKey)
	}

	return result
}

// Tags returns elb service tags.
func Tags(tags tftags.KeyValueTags) []*elb.Tag {
	result := make([]*elb.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &elb.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from elb service tags.
func KeyValueTags(ctx context.Context, tags []*elb.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// GetTagsIn returns elb service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*elb.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets elb service tags in Context.
func SetTagsOut(ctx context.Context, tags []*elb.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// UpdateTags updates elb service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.

func UpdateTags(ctx context.Context, conn elbiface.ELBAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &elb.RemoveTagsInput{
			LoadBalancerNames: aws.StringSlice([]string{identifier}),
			Tags:              TagKeys(removedTags.IgnoreAWS()),
		}

		_, err := conn.RemoveTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &elb.AddTagsInput{
			LoadBalancerNames: aws.StringSlice([]string{identifier}),
			Tags:              Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.AddTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).ELBConn(), identifier, oldTags, newTags)
}
