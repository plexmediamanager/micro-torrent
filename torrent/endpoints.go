package torrent

const (
    // Base API Endpoint
    ENDPOINT_API                            =   "%s/api/v2"

    // Authentication Endpoints
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#authentication
    ENDPOINT_AUTHENTICATION                 =   ENDPOINT_API + "/auth/"

    // [GET] Login
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#login
    ENDPOINT_AUTHENTICATION_LOGIN           =   ENDPOINT_AUTHENTICATION + "login"

    // [GET] Logout
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#logout
    ENDPOINT_AUTHENTICATION_LOGOUT          =   ENDPOINT_AUTHENTICATION + "logout"

    // Application endpoints
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#application
    ENDPOINT_APPLICATION                    =   ENDPOINT_API + "/app/"

    // [GET] Get application version
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-application-version
    ENDPOINT_APPLICATION_VERSION            =   ENDPOINT_APPLICATION + "version"

    // [GET] Get API version
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-api-version
    ENDPOINT_APPLICATION_API_VERSION        =   ENDPOINT_APPLICATION + "webapiVersion"

    // [GET] Get build info
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-build-info
    ENDPOINT_APPLICATION_BUILD_INFO         =   ENDPOINT_APPLICATION + "buildInfo"

    // [GET] Shutdown application
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#shutdown-application
    ENDPOINT_APPLICATION_SHUTDOWN           =   ENDPOINT_APPLICATION + "shutdown"

    // [GET] Get application preferences
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-application-preferences
    ENDPOINT_APPLICATION_PREFERENCES        =   ENDPOINT_APPLICATION + "preferences"

    // [POST] Set application preferences
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#set-application-preferences
    ENDPOINT_APPLICATION_SET_PREFERENCES    =   ENDPOINT_APPLICATION + "setPreferences"

    // [GET] Get default save path
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-default-save-path
    ENDPOINT_APPLICATION_DEFAULT_SAVE_PATH  =   ENDPOINT_APPLICATION + "defaultSavePath"

    // Log Endpoints
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#log
    ENDPOINT_LOG                            =   ENDPOINT_API + "/log/"

    // [GET] Get log
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-log
    ENDPOINT_LOG_GET                        =   ENDPOINT_LOG + "main"

    // [GET] Get peer log
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-peer-log
    ENDPOINT_LOG_PEERS                      =   ENDPOINT_LOG + "peers"

    // Sync Endpoints
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#sync
    ENDPOINT_SYNC                           =   ENDPOINT_API + "/sync/"

    // [GET] Get main data
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-main-data
    ENDPOINT_SYNC_MAIN                      =   ENDPOINT_SYNC + "maindata"

    // [GET] Get torrent peers data
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-torrent-peers-data
    ENDPOINT_SYNC_TORRENT_PEERS             =   ENDPOINT_SYNC + "torrentPeers"

    // Transfer info
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#transfer-info
    ENDPOINT_TRANSFER                       =   ENDPOINT_API + "transfer"

    // [GET] Get global transfer info
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-global-transfer-info
    ENDPOINT_TRANSFER_INFO                  =   ENDPOINT_TRANSFER + "info"

    // [GET] Get alternative speed limits state
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-alternative-speed-limits-state
    ENDPOINT_TRANSFER_ALT_SPEED_LIMITS      =   ENDPOINT_TRANSFER + "speedLimitsMode"

    // [GET] Toggle alternative speed limits
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#toggle-alternative-speed-limits
    ENDPOINT_TRANSFER_TOGGLE_SPEED_LIMITS   =   ENDPOINT_TRANSFER + "toggleSpeedLimitsMode"

    // [GET] Get global download limit
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-global-download-limit
    ENDPOINT_TRANSFER_GET_DOWNLOAD_LIMIT    =   ENDPOINT_TRANSFER + "downloadLimit"

    // [POST] Set global download limit
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#set-global-download-limit
    ENDPOINT_TRANSFER_SET_DOWNLOAD_LIMIT    =   ENDPOINT_TRANSFER + "setDownloadLimit"

    // [GET] Get global upload limit
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-global-upload-limit
    ENDPOINT_TRANSFER_GET_UPLOAD_LIMIT      =   ENDPOINT_TRANSFER + "uploadLimit"

    // [POST] Set global upload limit
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#set-global-upload-limit
    ENDPOINT_TRANSFER_SET_UPLOAD_LIMIT      =   ENDPOINT_TRANSFER + "setUploadLimit"

    // [ANY] Ban peers
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#ban-peers
    ENDPOINT_TRANSFER_BAN_PEERS             =   ENDPOINT_TRANSFER + "banPeers"

    // Torrent Management Endpoint
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#torrent-management
    ENDPOINT_TORRENT                        =   ENDPOINT_API + "/torrents/"

    // Get torrent list
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-torrent-list
    ENDPOINT_TORRENT_LIST                   =   ENDPOINT_TORRENT + "info"

    // [GET] Get torrent generic properties
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-torrent-generic-properties
    ENDPOINT_TORRENT_PROPERTIES             =   ENDPOINT_TORRENT + "properties"

    // [GET] Get torrent trackers
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-torrent-trackers
    ENDPOINT_TORRENT_TRACKERS               =   ENDPOINT_TORRENT + "trackers"

    // [GET] Get torrent web seeds
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-torrent-web-seeds
    ENDPOINT_TORRENT_WEB_SEEDS              =   ENDPOINT_TORRENT + "webseeds"

    // [GET] Get torrent contents
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-torrent-contents
    ENDPOINT_TORRENT_CONTENTS               =   ENDPOINT_TORRENT + "files"

    // [GET] Get torrent pieces' states
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-torrent-pieces-states
    ENDPOINT_TORRENT_PIECES_STATES          =   ENDPOINT_TORRENT + "pieceStates"

    // [GET] Get torrent pieces' hashes
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-torrent-pieces-hashes
    ENDPOINT_TORRENT_PIECES_HASHES          =   ENDPOINT_TORRENT + "pieceHashes"

    // [POST] Pause torrents
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#pause-torrents
    ENDPOINT_TORRENT_PAUSE                  =   ENDPOINT_TORRENT + "pause"

    // [POST] Resume torrents
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#resume-torrents
    ENDPOINT_TORRENT_RESUME                 =   ENDPOINT_TORRENT + "resume"

    // [POST] Delete torrents
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#delete-torrents
    ENDPOINT_TORRENT_DELETE                 =   ENDPOINT_TORRENT + "delete"

    // [POST] Recheck torrents
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#recheck-torrents
    ENDPOINT_TORRENT_RECHECK                =   ENDPOINT_TORRENT + "recheck"

    // [POST] Reannounce torrents
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#reannounce-torrents
    ENDPOINT_TORRENT_REANNOUNCE             =   ENDPOINT_TORRENT + "reannounce"

    // [POST] Add new torrent
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#add-new-torrent
    ENDPOINT_TORRENT_CREATE                 =   ENDPOINT_TORRENT + "add"

    // [POST] Add trackers to torrent
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#add-trackers-to-torrent
    ENDPOINT_TORRENT_ADD_TRACKERS           =   ENDPOINT_TORRENT + "addTrackers"

    // [POST] Edit trackers
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#edit-trackers
    ENDPOINT_TORRENT_EDIT_TRACKERS          =   ENDPOINT_TORRENT + "editTracker"

    // [POST] Remove trackers
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#remove-trackers
    ENDPOINT_TORRENT_REMOVE_TRACKERS        =   ENDPOINT_TORRENT + "removeTrackers"

    // [POST] Add peers
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#add-peers
    ENDPOINT_TORRENT_ADD_PEERS              =   ENDPOINT_TORRENT + "addPeers"

    // [POST] Increase torrent priority
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#increase-torrent-priority
    ENDPOINT_TORRENT_INCREASE_PRIORITY      =   ENDPOINT_TORRENT + "increasePrio"

    // [POST] Decrease torrent priority
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#decrease-torrent-priority
    ENDPOINT_TORRENT_DECREASE_PRIORITY      =   ENDPOINT_TORRENT + "decreasePrio"

    // [POST] Maximal torrent priority
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#maximal-torrent-priority
    ENDPOINT_TORRENT_MAX_PRIORITY           =   ENDPOINT_TORRENT + "topPrio"

    // [POST] Minimal torrent priority
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#minimal-torrent-priority
    ENDPOINT_TORRENT_MIN_PRIORITY           =   ENDPOINT_TORRENT + "bottomPrio"

    // [POST] Set file priority
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#set-file-priority
    ENDPOINT_TORRENT_SET_FILE_PRIORITY      =   ENDPOINT_TORRENT + "filePrio"

    // [POST] Get torrent download limit
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#get-torrent-download-limit
    ENDPOINT_TORRENT_GET_DOWNLOAD_LIMIT     =   ENDPOINT_TORRENT + "downloadLimit"

    // [POST] Set torrent download limit
    // https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation#set-torrent-download-limit
    ENDPOINT_TORRENT_SET_DOWNLOAD_LIMIT     =   ENDPOINT_TORRENT + "setDownloadLimit"
)
