import requests
import json

function_name = "update"

source_code = '''
    def update(self) -> None:
        """Run function to get device details."""
        self.get_details()
        '''
        
context = '''
    def get_details(self) -> None:
        """Build Air Purifier details dictionary."""
        body = Helpers.req_body(self.manager, 'devicedetail')
        body['uuid'] = self.uuid
        head = Helpers.req_headers(self.manager)

        r, _ = Helpers.call_api(
            '/131airPurifier/v1/device/deviceDetail',
            method='post',
            headers=head,
            json_object=body,
        )

        if r is not None and Helpers.code_check(r):
            self.device_status = r.get('deviceStatus', 'unknown')
            self.connection_status = r.get('connectionStatus', 'unknown')
            self.details['active_time'] = r.get('activeTime', 0)
            self.details['filter_life'] = r.get('filterLife', {})
            self.details['screen_status'] = r.get('screenStatus', 'unknown')
            self.mode = r.get('mode', self.mode)
            self.details['level'] = r.get('level', 0)
            self.details['air_quality'] = r.get('airQuality', 'unknown')
        else:
            logger.debug('Error getting %s details', self.device_name)
        '''


pre_answers = '''
    get_details:
    {'device_status': 'deviceStatus', 'connection_status': 'connectionStatus', 'active_time': 'activeTime', 'filter_life': 'filterLife', 'screen_status': 'screenStatus', 'mode': 'mode', 'level': 'level', 'air_quality': 'airQuality'}

    active_time:
    {'Type': 'Read', 'active_time': 'activeTime'}

    fan_level:
    {'Type': 'Read', 'level': None}

    filter_life:
    {
        "Type": "Read",
        "filter_life": "filterLife"
    }

    air_quality:
    {
        "Type": "Read",
        "air_quality": None
    }

    screen_status:
    {
        "Type": "Read",
        "screen_status": None
    }

    toggle_display:
    {
        "Type": "Write",
        "screen_status": "status"
    }

    turn_on:
    {
        "Type": "Write",
        "device_status": "status"
    }

    turn_off:
    {
        "Type": "Write",
        "device_status": "off"
    }

    change_fan_speed:
    {
        "Type": "Write",
        "level": "speed"
    }

    mode_toggle:
    {
        "Type": "Write",
        "mode": "mode",
        "level": "level"
    }

    turn_on_display:
    {
        "Type": "Write",
        "screen_status": "on"
    }

    turn_off_display:
    {
        "Type": "Write",
        "screen_status": "off"
    }

    auto_mode:
    {
        "Type": "Write",
        "mode": "auto"
    }

    manual_mode:
    {
        "Type": "Write",
        "mode": "manual",
        "level": 1
    }

    sleep_mode:
    {
        "Type": "Write",
        "mode": "sleep",
        "level": 1
    }
        '''

payload = json.dumps({
   "model": "gpt-3.5-turbo",
   "messages": [
        {
            "role": "system",
            "content": "The first function is an SDK API codes for a fan. Others functions or classes provide context related to the first one. Consider the semantics and functionality of the first function to determine if it is a read method or a write method. If it is a write method, list all the device properties it modifies. If any of the API's parameters modify these device properties, list their relationships. If it is a read method, specify the device properties it retrieves. Provide the response in JSON format without extra information or explanation., structured as: {'Type': 'Write/Read', 'key': 'value/None', 'key': 'value/None' ... }. Each key is the device property you retrieved from the first function and its data type, and the value is the parameter of the first function that corresponds to the device property. If a device property does not have a corresponding parameter, the value should be None. Remember, there may be cases where no device properties are available. In such cases, no device properties need to be returned. At the same time, we also provide the relationship of device properties from previous rounds of answers. If the device properties corresponding to this API have appeared in previous answers, please ensure that the device property names are consistent with the previous answers. "
        },
        {
            "role": "user",
            "content": source_code
        },
        {
            "role": "user",
            "content": context
        },
        {
            "role": "user",
            "content": pre_answers
        },
   ]
})

headers = {
   'Accept': 'application/json',
   'Authorization': 'XXX',
   'User-Agent': 'XXX',
   'Content-Type': 'application/json'
}

url = "XXX"

response = requests.request("POST", url, headers=headers, data=payload)
response_content = response.json()["choices"][0]["message"]["content"]

