import http from "k6/http";
import { check, sleep } from "k6";
import { Trend } from "k6/metrics";

let myTrend = new Trend("my_custom_trend");

export const options = {
  scenarios: {
    constant_load: {
      executor: "constant-vus",
      vus: 10, // 가상 사용자 수
      duration: "30s", // 테스트 지속 시간
    },
    ramping_load: {
      executor: "ramping-vus",
      startVUs: 0,
      stages: [
        { duration: "10s", target: 20 }, // 10초 동안 가상 사용자 수를 20으로 증가
        { duration: "10s", target: 20 }, // 10초 동안 가상 사용자 수를 유지
        { duration: "10s", target: 0 }, // 10초 동안 가상 사용자 수를 0으로 감소
      ],
    },
    spike_test: {
      executor: "ramping-arrival-rate",
      startRate: 10, // 시작 요청 수 (초당 10개)
      stages: [
        { duration: "10s", target: 100 }, // 10초 동안 초당 요청 수를 100으로 증가
        { duration: "10s", target: 100 }, // 10초 동안 초당 요청 수를 유지
        { duration: "10s", target: 10 }, // 10초 동안 초당 요청 수를 10으로 감소
      ],
      preAllocatedVUs: 50, // 미리 할당된 가상 사용자 수
      maxVUs: 100, // 최대 가상 사용자 수
    },
  },
};

export default function () {
  let res = http.get("http://localhost:3000");

  check(res, {
    "status is 200": (r) => r.status === 200,
  });

  // 사용자 정의 메트릭 기록
  myTrend.add(res.timings.duration);

  // 각 가상 사용자 사이에 대기 시간 추가
  sleep(1);
}
