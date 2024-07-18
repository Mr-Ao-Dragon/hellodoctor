import {List} from "immutable";
import axios from 'axios';

const apiBaseUrl = "https://" + import.meta.env.VITE_API_HOSTNAME + '/' + import.meta.env.VITE_API_VERSION;
let apiResponse;

function list() {
    const url = apiBaseUrl + "/doctor/index/list";
    let config = {
        method: 'get',
        url: url,
        headers: {
            'User-Agent': 'Apifox/1.0.0 (https://apifox.com)'
        }
    };
    axios(config)
        .then(function (response) {
            if (response.data.code === 200)
                apiResponse = List(response.data.data);
            else {
                console.log(response.data.msg);
            }
        })
        .catch(function (error) {
            console.log(error);
        });
    return apiResponse;
}

function oneInfo(id) {
    const url = apiBaseUrl + "/doctor/one_info"
    let request = {
        id: id
    }
    let config = {
        method: 'post',
        url: url,
        headers: {
            'User-Agent': 'Apifox/1.0.0 (https://apifox.com)',
            'Content-Type': 'application/json'
        },
        data: JSON.stringify(request)
    };
    axios(config)
        .then(function (response) {
            if (response.data.code === 200)
                apiResponse = response.data.data;
            else {
                console.log(response.data.msg);
            }
        })
        .catch(function (error) {
            console.log(error);
        });
    return apiResponse;
}
