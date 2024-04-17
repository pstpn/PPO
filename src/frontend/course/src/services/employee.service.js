import axios from 'axios';
import authHeader from './auth-header';

const API_URL = 'http://localhost:8081/api/v1/';

class UserService {
    // getHomePage() {
    //     return axios.get(API_URL + 'home');
    // }

    createEmployeeInfoCard() {
        return axios.post(API_URL + 'infoCards', { headers: authHeader() });
    }

    getEmployeeInfoCard() {
        return axios.get(API_URL + 'infoCards', { headers: authHeader() });
    }

    getAdminBoard() {
        return axios.get(API_URL + 'admin', { headers: authHeader() });
    }
}

export default new UserService();