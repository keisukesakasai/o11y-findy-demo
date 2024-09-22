import { sleep } from 'k6';
import http from 'k6/http';

const baseEndpoints = ['http://app-service-v1:8080', 'http://app-service-v2:8080'];

export const options = {
    vus: 1, // 1 virtual user
    duration: '0s', // Run indefinitely
};

let endpointIndex = 0; // Track which endpoint to use

export default function () {
    // Send a request to app-service-v1 or app-service-v2 alternately
    const baseEndpoint = baseEndpoints[endpointIndex];
    http.get(baseEndpoint);

    // Switch between app-service-v1 and app-service-v2 for the next request
    endpointIndex = (endpointIndex + 1) % baseEndpoints.length;

    // Wait for 1 second before sending the next request
    sleep(1);
}
