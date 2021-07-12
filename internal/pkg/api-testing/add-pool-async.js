import http from 'k6/http';

export let options = {
    vus: 200,
    duration: '1m',
};

export function setup() {
    var url = 'http://localhost:8000/v1/pool/add';

    var values = []
    for (let i = 0; i < 1000000; i++) {
        values.push(Math.floor(Math.random() * 1000))
    }
    var payload = JSON.stringify({
        pool_id: 1,
        pool_values: values,
    });

    var params = {
        headers: {
            'Content-Type': 'application/json',
        },
        timeout: "10m",
    };

    http.post(url, payload, params);
}

export default function () {
    var url = 'http://localhost:8000/v1/pool/add';

    var values = []
    for (let i = 0; i < 100; i++) {
        values.push(Math.floor(Math.random() * 10000))
    }
    var payload = JSON.stringify({
        pool_id: 1,
        pool_values: values,
    });

    var params = {
        headers: {
            'Content-Type': 'application/json',
        },
        timeout: "10m",
    };

    http.post(url, payload, params);
}

