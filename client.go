package opensearch

type ClientOptions struct {
     Addresses []string
     Insecure bool
     Username string
     Password string
     AWSCredentialsURI string
}

// func NewClient(ctx context.Context, opts *ClientOptions) 