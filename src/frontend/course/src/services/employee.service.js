import axios from 'axios';
import authHeader from './auth-header';

const API_URL = 'http://localhost:8081/';

class UserService {
    createEmployeeInfoCard() {
        return axios.post(API_URL + 'infocards', {}, {headers: authHeader()}).then(
            response => {
                return response.data;
            });
    }

    getEmployeeInfoCard() {
        // return axios.get(API_URL + 'infocards' + '/' + , { headers: authHeader() });
    }

    getAdminBoard() {
        return axios.get(API_URL + 'admin', { headers: authHeader() });
    }
}

export default new UserService();