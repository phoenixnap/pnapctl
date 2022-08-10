# Generates mocks for all APIs.
PREFIX="make generate-mock"
MOCK_SOURCE="MOCK_SOURCE=common/client/"
CLIENT="/client.go"
MOCK_DESTINATION="MOCK_DESTINATION=testsupport/mocks/sdkmocks/mock_"
MOCK_CLIENT="_client.go"
PACKAGE="MOCK_PACKAGE=sdkmocks"

is=(
    "bmcapi"
    "audit"
    "ip"
    "networks"
    "rancher"
    "tags"
)

for api in "${is[@]}"; do
    echo -e "Running $PREFIX ${MOCK_SOURCE}${api}${CLIENT} ${MOCK_DESTINATION}${api}${MOCK_CLIENT} $PACKAGE..." 
    $PREFIX ${MOCK_SOURCE}${api}${CLIENT} ${MOCK_DESTINATION}${api}${MOCK_CLIENT} $PACKAGE
done