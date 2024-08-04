import json

source_code = ""
context = ""
pre_outputs = ""

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
            "content": pre_outputs
        },
   ]
})