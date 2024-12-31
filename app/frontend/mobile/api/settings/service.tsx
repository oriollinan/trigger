import { Env } from "@/lib/env";
import AsyncStorage from "@react-native-async-storage/async-storage";

export class SettingsService {
    static async getBaseUrl() {
        return `${Env.NGROK}/api/settings`;
    }

    static async getSettings() {
        try {
            const baseUrl = await this.getBaseUrl();
            const response = await fetch (`${baseUrl}/me`, {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${await AsyncStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                }
            });
            if (response.status !== 200) {
                console.log('get settings failed', response.status);
                throw new Error('Something went wrong.');
            }
            const data = await response.json();
            return data;
        } catch (error) {
            console.error("Catched Get Settings Error:", error);
            throw error;
        }
    }
}
