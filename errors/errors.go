
package errors

import "github.com/plexmediamanager/service/errors"

const (
    ServiceID       errors.Service      =   5
)

var (
    RequestCreationError = errors.Error{
        Code:               errors.Code{
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    1,
        },
        Message:            "Failed to create new HTTP Request",
    }
    StringToIntError = errors.Error{
        Code:               errors.Code{
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    2,
        },
        Message:            "Failed to convert string to integer",
    }
    UnmarshalError = errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    3,
        },
        Message:    "Failed to convert bytes to structure",
    }
    // Network Errors
    RequestCouldNotBeCompleted = errors.Error{
        Code:               errors.Code{
            Service:        ServiceID,
            ErrorType:      errors.TypeNetwork,
            ErrorNumber:    1,
        },
        Message:            "Failed to perform request to: %s",
    }
    TorrentAuthenticationError = errors.Error{
        Code:               errors.Code{
            Service:        ServiceID,
            ErrorType:      errors.TypeNetwork,
            ErrorNumber:    2,
        },
        Message:            "Failed to log in with the given credentials",
    }
)