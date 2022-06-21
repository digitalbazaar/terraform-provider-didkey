package didkey

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/multiformats/go-multibase"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
)

func resourceDidKey() *schema.Resource {
	return &schema.Resource{
		Description: `
The resource ` + "`random_id`" + ` generates random numbers that are intended to be
used as unique identifiers for other resources.
This resource *does* use a cryptographic random number generator in order
to minimize the chance of collisions, making the results of this resource
when a 16-byte identifier is requested of equivalent uniqueness to a
type-4 UUID.
This resource can be used in conjunction with resources that have
the ` + "`create_before_destroy`" + ` lifecycle flag set to avoid conflicts with
unique names during the brief period where both the old and new resources
exist concurrently.
`,
		CreateContext: CreateID,
		ReadContext:   NoOp,
		DeleteContext: RemoveResourceFromState,
		Importer: &schema.ResourceImporter{
			StateContext: ImportID,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The generated id presented in base64.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"keepers": {
				Description: "Arbitrary map of values that, when changed, will trigger recreation of " +
					"resource. See [the main provider documentation](../index.html) for more information.",
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},

			"secret_seed_multibase": {
				Description: "The private seed.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"public_did": {
				Description: "The public key.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func CreateID(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	byteLength := 8
	bytes := make([]byte, byteLength)

	n, err := rand.Reader.Read(bytes)
	if n != byteLength {
		return append(diags, diag.Errorf("generated insufficient random bytes: %s", err)...)
	}
	if err != nil {
		return append(diags, diag.Errorf("error generating random bytes: %s", err)...)
	}

	b64Str := base64.RawURLEncoding.EncodeToString(bytes)
	d.SetId(b64Str)

	public, private, _ := ed25519.GenerateKey(rand.Reader)
	private_seed := private.Seed()

	code := multicodec.Ed25519Pub
	multicodecPublicBytes := append([]byte{byte(code), 0x01}, public...)
	b58MultibasePublic, _ := multibase.Encode(multibase.Base58BTC, multicodecPublicBytes)

	b58MultihashPrivateSeed, _ := multihash.Encode(private_seed, multihash.IDENTITY)
	b58MultibasePrivateSeed, _ := multibase.Encode(multibase.Base58BTC, b58MultihashPrivateSeed)

	if err := d.Set("public_did", "did:key:"+b58MultibasePublic); err != nil {
		return append(diags, diag.Errorf("error setting public_did: %s", err)...)
	}
	if err := d.Set("secret_seed_multibase", b58MultibasePrivateSeed); err != nil {
		return append(diags, diag.Errorf("error setting secret_seed_multibase: %s", err)...)
	}

	return diags
}

func NoOp(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func ImportID(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}
