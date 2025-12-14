import axiosConfig from "./axiosConfig.js";
import axios from "axios";

export const fetchAuthInfo = async () => {
    try {
        return await axiosConfig.get("/auth/auth-info");
    } catch (error) {
        console.error("Error fetching auth info:", error);
        throw error;
    }
};

export const authenticateUser = async (code) => {
    try {
        const baseUrl = axiosConfig.defaults.baseURL;
        let url = baseUrl + `/auth/authenticate?code=${code}`;
        return await axios.get(url);
    } catch (error) {
        console.error("Error authenticating user:", error);
        throw error;
    }
};
