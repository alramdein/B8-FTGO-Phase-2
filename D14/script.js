import http from 'k6/http';
import { check } from 'k6';

export const options = {
  vus: 100,
  duration: '5s',
};

export default function() {
  const res = http.get('http://localhost:8080/users');

  check(res, {
     'status was 200': (r) => r.status == 200 
  });
}

// gimana kalo --> IP Address + 1 endpoint == rate limit
// go concunreey --> op