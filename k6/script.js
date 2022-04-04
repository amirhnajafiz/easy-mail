import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
    const url = 'http://127.0.0.1:5000/mail/send';
    const payload = JSON.stringify({
        from: "najafizadeh21@gmail.com",
        to: "najafizadeh21@gmail.com",
        subject: "Easy mail",
        text: "This is a test message.",
        validation: "no"
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
    sleep(1);
}
