import axios from 'axios';

const API_URL = 'http://localhost:8081/';

class AuthService {
    login(user) {
        return axios
            .post(API_URL + 'login', {
                phoneNumber: user.phoneNumber,
                password: user.password,
            })
            .then(response => {
                if (response.data) {
                    localStorage.setItem('user', JSON.stringify(response.data));
                }

                return response.data
            });
    }

    logout() {
        localStorage.removeItem('user');
    }

    register(user) {
        return axios.post(API_URL + 'register', {
            phoneNumber: user.phoneNumber,
            name: user.name,
            surname: user.surname,
            companyID: user.selectedCompany + 1,
            post: user.post,
            password: user.password,
            dateOfBirth: user.dateOfBirth,
        })
        .then(response => {
            if (response.data) {
                localStorage.setItem('user', JSON.stringify(response.data));
            }

            return response.data;
        });
    }

    refreshTokens(user) {
        return axios.post(API_URL + 'refresh', {
            accessToken: user.accessToken,
            refreshToken: user.refreshToken,
        })
        .then(response => {
            if (response.data) {
                localStorage.setItem('user', JSON.stringify(response.data));
            }

            return response.data;
        })
    }
}

export default new AuthService();