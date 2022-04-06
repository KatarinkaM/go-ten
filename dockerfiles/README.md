# Obscuro enclave service Docker image

To create the Docker image for an Obscuro enclave service running in SGX simulation mode:

    docker build -t obscuro_enclave - < dockerfiles/enclave.Dockerfile

This is a required step before running the Docker integration tests.

To run the image as a container, where `XXX` is the port on which to expose the enclave service's RPC endpoints on the 
local machine, and `YYY` is the address of the node that this enclave service is for as an integer:

    docker run -p XXX:11000/tcp obscuro_enclave --nodeID YYY --address :11000