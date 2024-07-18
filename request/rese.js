import {List} from "immutable";
import axios from 'axios';

const apiBaseUrl = "https://" + import.meta.env.VITE_API_HOSTNAME + '/' + import.meta.env.VITE_API_VERSION;

function getOrderList(id, token) {
    var axios = require('axios');
    var config = {
        method: 'get', url: apiBaseUrl + '/rese/info', headers: {
            'User-Agent': 'Apifox/1.0.0 (https://apifox.com)'
        },
        params: {
            id: id, token: token
        }
    };
    axios(config)
        .then(function (response) {
            console.log(JSON.stringify(response.data));
        })
        .catch(function (error) {
            console.log(error);
        });
}
