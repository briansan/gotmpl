# gotmpl
Minimal implementation of text rendering using Go templates

# Usage
```
echo "tag: 1.17.7" > args.yaml
docker run -v $(pwd):/var briansan/gomtpl
```
By default, gotmpl renders text off the template found at /var/manifest.gotmpl using the YAML defintion in /var/args.yaml.
