#! /bin/sh
# Script to do the following
# 1. Run docker-compose to spin off dynamodb locally on localhost:8000
# 2. Create the following tables:
#   a) UserDetails - Store user details
#   b) Messages - Store the messages sent and received

echo "Starting up dynamodb, please make sure docker daemon is up and running"
exec docker-compose restart &
BACK_PID=$!
while kill -0 $BACK_PID; do
    echo "Dynaomodb is still starting..."
    sleep 1
done
echo "Creating tables now, please make sure aws cli is installed, https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-mac.html"
aws dynamodb create-table \
                  --table-name UserDetails \
                  --attribute-definitions \
                    AttributeName=UserId,AttributeType=S \
                    AttributeName=Email,AttributeType=S \
                 --key-schema \
                    AttributeName=UserId,KeyType=HASH \
                    AttributeName=Email,KeyType=RANGE \
                 --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
                 --endpoint-url http://localhost:8000 \
                 --global-secondary-indexes \
                 "[{\"IndexName\": \"EmailIndex\",
	                    \"KeySchema\":[{\"AttributeName\":\"Email\",\"KeyType\":\"HASH\"}],
	                    \"Projection\":{\"ProjectionType\":\"ALL\"},
	                    \"ProvisionedThroughput\":{\"ReadCapacityUnits\":1, \"WriteCapacityUnits\":1}},
	                    {\"IndexName\": \"UserIdIndex\",
	                    \"KeySchema\":[{\"AttributeName\":\"UserId\",\"KeyType\":\"HASH\"}],
	                    \"Projection\":{\"ProjectionType\":\"ALL\"},
	                    \"ProvisionedThroughput\":{\"ReadCapacityUnits\":1, \"WriteCapacityUnits\":1}}]"

aws dynamodb create-table \
                  --table-name MessageDetails \
                  --attribute-definitions \
                    AttributeName=ConversationId,AttributeType=S \
                    AttributeName=Sent,AttributeType=S \
                 --key-schema \
                    AttributeName=ConversationId,KeyType=HASH \
                    AttributeName=Sent,KeyType=RANGE \
                 --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
                 --endpoint-url http://localhost:8000 \
                 --global-secondary-indexes \
                 "[{\"IndexName\": \"ConversationIdIndex\",
	                    \"KeySchema\":[{\"AttributeName\":\"ConversationId\",\"KeyType\":\"HASH\"},
	                    {\"AttributeName\":\"Sent\",\"KeyType\":\"RANGE\"}],
	                    \"Projection\":{\"ProjectionType\":\"ALL\"},
	                    \"ProvisionedThroughput\":{\"ReadCapacityUnits\":1, \"WriteCapacityUnits\":1}}]"


echo "Following tables were created :"
aws dynamodb list-tables --endpoint-url http://localhost:8000

echo "Creating Message Queue"
aws sqs create-queue --queue-name Messages --endpoint-url http://localhost:9324