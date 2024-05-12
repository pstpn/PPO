import EmployeeService from "../services/employee.service";

export const employee = {
    namespaced: true,
    state: {
        message: "",
    },
    actions: {
        fillProfile({ commit }, formData) {
            return EmployeeService.fillProfile(formData).then(
                infoCard => {
                    return Promise.resolve(infoCard);
                },
                error => {
                    return Promise.reject(error);
                })
        },
        getProfile({ commit }) {
            return EmployeeService.getProfile().then(
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
