VARFILE=$(dirname "$0")/.vars

if [ -f "$VARFILE" ]; then
    while IFS='|', read -r key value env; do
        if $env = "true"; then
            eval export $key=\$value
        else
            eval $key=\$value
        fi
        echo "$key | $value | $env"
    done < $VARFILE
fi