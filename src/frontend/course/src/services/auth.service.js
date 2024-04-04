import axios, {defaults} from 'axios';

const API_URL = 'http://localhost:8081/api/v1/';

// FIXME
axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
class AuthService {
    login(user) {
        return axios
            .post(API_URL + 'signin', {
                phoneNumber: user.phoneNumber,
                password: user.password
            })
            .then(response => {
                if (response.data.accessToken) {
                    localStorage.setItem('user', JSON.stringify(response.data));
                }

                return response.data;
            });
    }

    logout() {
        localStorage.removeItem('user');
    }

    // FIXME
    register(user) {
        return axios.post(API_URL + 'signup', {
            phoneNumber: user.phoneNumber,
            password: user.password
        }, {
            headers: {
                'Content-Type': 'application/json',
            }
        });
    }
}

export default new AuthService();