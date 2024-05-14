import axios from 'axios';
import authHeader from './auth-header';

const API_URL = 'http://localhost:8081/';

class EmployeeService {
    fillProfile(formData) {
        return axios.post(API_URL + 'profile', formData, { headers: { ...authHeader(), 'Content-Type': 'multipart/form-data' } })
            .then(response => {
                return response.data;
            });
    }

    getProfile() {
        return axios.get(API_URL + 'profile', { headers: authHeader() })
            .then(response => {
                return response.data
            });
    }

    getEmployeePhoto() {
        return axios.get(API_URL + 'employee-photo', { headers: authHeader(), responseType: "arraybuffer" })
            .then(response => {
                return URL.createObjectURL(new Blob([response.data], { type: 'image/jpeg' }));
            });
    }
}

export default new EmployeeService();