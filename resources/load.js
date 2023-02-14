import http from 'k6/http';
import { check, group, sleep } from "k6";
import { Rate } from "k6/metrics";

new Rate("check_failure_rate");

export const options = {
//Stress Testing is a type of load testing used to determine the limits of the system.
// The purpose of this test is to verify the stability and reliability of the system under extreme conditions.
  stages: [

    { duration: '2m', target: 100 }, // below normal load

    { duration: '5m', target: 100 },

    { duration: '2m', target: 100 }, // normal load

    { duration: '5m', target: 100 },

    { duration: '2m', target: 100 }, // around the breaking point

    { duration: '5m', target: 100 },

    { duration: '2m', target: 100 }, // beyond the breaking point

    { duration: '5m', target: 100 },

    { duration: '10m', target: 0 }, // scale down. Recovery stage.

  ],


  insecureSkipTLSVerify: true,
  noConnectionReuse: false,
  vus: 100,
  thresholds: {
    // 99% of requests must complete below 1.5s
    "http_req_duration": ["p(99)<500"],
    "check_failure_rate": ["rate<0.01"],
    "http_reqs": ["rate<=600"],
  }
};

export default function () {
  http.get('http://localhost:8081/api/employees');
  sleep(1);
}