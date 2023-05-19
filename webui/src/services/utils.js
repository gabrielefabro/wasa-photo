var utils = {}

/*
    *  Convert an error to a string. If the error has a response, the response data is returned.
    *  Otherwise, the error is converted to a string.
    *  @param {Error} error - The error to convert.
    *  @returns {String} The string representation of the error.
*/
utils.errorToString = (error) => {
    console.log(error);
    if (error.hasOwnProperty('response')) {
        return error.response.data;
    } else {
        return error.toString();
    }
}

export default utils;