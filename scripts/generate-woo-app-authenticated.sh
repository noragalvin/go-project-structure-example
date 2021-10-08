# Read more: https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#rest-api-keys
# Bash example
STORE_URL='http://localhost:8080'
ENDPOINT='/wc-auth/v1/authorize'
PARAMS="app_name=EcomSolid Headless Storefront&scope=read_write&user_id=123&return_url=http://localhost:3000/helloworld&callback_url=https://norahaha.ap.ngrok.io/callback-endpoint"
QUERY_STRING="$(perl -MURI::Escape -e 'print uri_escape($ARGV[0]);' "$PARAMS")"
QUERY_STRING=$(echo $QUERY_STRING | sed -e "s/%20/\+/g" -e "s/%3D/\=/g" -e "s/%26/\&/g")

echo "$STORE_URL$ENDPOINT?$QUERY_STRING"
