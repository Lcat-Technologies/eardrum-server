package errors

import (
    "net/http"
    
    // Assuming this import path for your central error struct
    customErrors "github.com/GigaDesk/eardrum-interfaces/errors" 
)

// --- 4XX CLIENT/BUSINESS LOGIC ERRORS ---

// NewUnauthorizedError returns a 401 Unauthorized error.
// Used when credentials provided are invalid or missing (e.g., token check).
func NewUnauthorizedError(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusUnauthorized, 
        message, 
        nil, 
    )
}

// NewBadRequestError returns a 400 bad request error.
// Used when input provided doesnt match the required criteria(e.g setting a weak password)
func NewBadRequestError(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusBadRequest, 
        message, 
        nil, 
    )
}


// NewNotFound returns a 404 not found error.
// Used when a record needed in an operation does not exist
func NewNotFoundError(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusNotFound, 
        message, 
        nil, 
    )
}

// NewForbiddenError returns a 403 forbidden error.
// Used when resource exists but not allowed to act
func NewForbiddenError(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusForbidden, 
        message, 
        nil, 
    )
}

// ErrDBPersistenceFailure returns a 500 Internal Server Error.
// Used for unexpected errors (deadlocks etc.).
func ErrPersistenceFailure(message string) *customErrors.PublicError {
    return customErrors.NewHTTPError(
        http.StatusInternalServerError, 
        message,
        nil,
    )
}
