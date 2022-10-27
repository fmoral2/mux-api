import http from 'k6/http';
import { check, group, sleep } from "k6";
import { Rate } from "k6/metrics";

new Rate("check_failure_rate");

export let options = {

  insecureSkipTLSVerify: true,
  noConnectionReuse: false,
  vus: 100,
  duration: '10s',
  iterations: 1000,
  thresholds: {
    "http_req_duration": ["p(95)<500"],
    "check_failure_rate": [ "rate<0.01"]
  }
};
export default function () {

  http.get('http://localhost:8081/api/employees');
  sleep(1);

}