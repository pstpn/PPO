import EmployeeService from "../services/employee.service";
import AuthService from "@/services/auth.service";

const employeeData = JSON.parse(localStorage.getItem('employeeData'));

export const employee = {
    state: {
        createdInfoCard: false,
        message: "",
        employeeInfo: null,
    },
    actions: {
        createEmployeeInfoCard({ commit }, employeeData) {
            return EmployeeService.createEmployeeInfoCard(employeeData).then(
                infoCard => {
                    commit('createdInfoCard', true);
                    commit('employeeInfo', infoCard);
                    console.log(employeeData);
                    console.log(infoCard);
                    return Promise.resolve(infoCard);
                },
                error => {
                    commit('createdInfoCard', false);
                    commit('employeeInfo', null);
                    return Promise.reject(error);
                }
            );
        },
        getEmployeeInfoCard({ commit }) {
            commit("setLoading", true);
            try {
                const response = EmployeeService.getEmployeeInfoCard();
                commit("setLoading", false);
                commit("setEmployeeInfo", response.data);
                return response.data;
            } catch (error) {
                commit("setLoading", false);
                commit("setMessage", error.message);
                throw error;
            }
        },
    },
    mutations: {
        setCreated(state, value) {
            state.createdInfoCard = value;
        },
        setLoading(state, value) {
            state.loading = value;
        },
        setMessage(state, message) {
            state.message = message;
        },
        setEmployeeInfo(state, info) {
            state.employeeInfo = info;
        },
    },
};
