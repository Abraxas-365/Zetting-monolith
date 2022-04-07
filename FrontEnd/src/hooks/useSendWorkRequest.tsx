import { CommonActions } from "@react-navigation/native";
import { apiCalls } from "../api/apiCalls";
import { WorkRequest } from "../interfaces/appInterfaces";

export const sendWorkRequest = async (worker_id: any, project_id: any, token: any, navigation: any) => {

    const config = {
        headers: { Authorization: `Bearer ${token}` }
    };
    console.log(config);

    try {
        const { data } = await apiCalls.post<WorkRequest>('/api/work-request/new', { worker_id, project_id }, config);
        await apiCalls.post('/api/notification/new/type=new-work-request', { "notified_id": worker_id, "reference_id": data.id }, config)
        navigation.dispatch(
            CommonActions.reset({
                index: 0,
                routes: [{ name: 'BottomTabs' }
                ]

            }))

    } catch (err) {
        console.error(err);
    }
}
