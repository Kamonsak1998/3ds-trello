import { mainAPIOptions } from "../constant/config";
import { BaseService } from "./BaseService";

export class BoardService extends BaseService {
    fetchDashboard = () => {
        return this.requester.get('/getdashboard', mainAPIOptions())
    }

    fetchHistory = ({ idBoard }) => {
        return this.requester.post('/gethistory', { idBoard }, mainAPIOptions())
    }

    fetchLeaderboard = ({ idBoard }) => {
        return this.requester.post('/getleaderboard', { idBoard }, mainAPIOptions())
    }

    // fetchSetdatetime = ({
    //     startDate,
    //     sprintDay,
    //     endDate,
    //     idBoard,
    //     boardName
    // }) => {
    //     return this.requester.post('/setdate', {
    //         startDate,
    //         sprintDay,
    //         endDate,
    //         idBoard,
    //         boardName
    //     }, mainAPIOptions())
    // }

    fetchchecksetdate = ({ idBoard }) => {
        return this.requester.post('/checksetdate', { idBoard }, mainAPIOptions())
    }

    fetchsetmember = ({ idBoard }) => {
        return this.requester.post('/setmember', { idBoard }, mainAPIOptions())
    }

    // fetchcheckscoresize = () => {
    //     return this.requester.get('/checkscoresize', mainAPIOptions())
    // }

    // fetchsetscoresize = ({ Points }) => {
    //     return this.requester.post('/setscoresize', { Points }, mainAPIOptions())
    // }

    fetchsettingdata = ({
        sprintDate, scoreSize,selectList
    }) => {
        return this.requester.post('/settingdata', { sprintDate, scoreSize,selectList }, mainAPIOptions())
    }

    fetchchecksetting = ({ idBoard }) => {
        return this.requester.post('/checksetting', { idBoard }, mainAPIOptions())
    }

}