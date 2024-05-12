import axios from 'axios';
import authHeader from './auth-header';

const API_URL = 'http://localhost:8081/';

class UserService {
    fillProfile(formData) {
        return axios.post(API_URL + 'profile', formData, { headers: { ...authHeader(), 'Content-Type': 'multipart/form-data' } })
            .then(response => {
                return response.data;
            });
    }

    getProfile() {
        // return axios.get(API_URL + 'profile' + '/' + , { headers: authHeader() });
    }
}

export default new UserService();