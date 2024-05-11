import EmployeeService from "../services/employee.service";

export const employee = {
    namespaced: true,
    state: {
        message: "",
    },
    actions: {
        createEmployeeInfoCard({ commit }) {
            return EmployeeService.createEmployeeInfoCard().then(
                infoCard => {
                    return Promise.resolve(infoCard);
                },
                error => {
                    return Promise.reject(error);
                })
        },
        getEmployeeInfoCard({ commit }) {
            return EmployeeService.getEmployeeInfoCard().then(
                infoCard => {
                    return Promise.resolve(infoCard);
                },
                error => {
                    return Promise.reject(error);
                })
        }
    },
    mutations: {
        setLoading(state, value) {
            state.loading = value;
        },
        setMessage(state, message) {
            state.message = message;
        },
    },
};
