#!/bin/sh -xe -o pipefail

base_dir="`dirname $0`"

curl_post_json () {
  curl --data-binary @"$base_dir/$1" \
    -X POST \
    -H 'content-type: application/json; charset=UTF-8' \
    -v \
    "http://localhost:3000/?webhook=$SLACK_WEBHOOK"
}

curl_post_json issue_created.json

curl_post_json comment_created.json
curl_post_json issue_updated_commented.json

curl_post_json issue_updated_summary.json

curl_post_json issue_updated_assigned.json

curl_post_json issue_updated_transition.json

curl_post_json comment_deleted.json
curl_post_json issue_updated_comment_deleted.json

curl_post_json issue_deleted.json
