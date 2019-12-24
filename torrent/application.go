package torrent

import (
    "github.com/plexmediamanager/micro-torrent/errors"
    "strconv"
    "strings"
)

type ApplicationVersion struct {
    Major                               uint64
    Minor                               uint64
    Patch                               uint64
}

type ApplicationAPIVersion struct {
    Major                               uint64
    Minor                               uint64
}

type ApplicationBuildInfo struct {
    Bitness                             uint64              `json:"bitness"`
    Boost                               string              `json:"boost"`
    LibTorrent                          string              `json:"libtorrent"`
    OpenSSL                             string              `json:"openssl"`
    QTVersion                           string              `json:"qt"`
    ZLib                                string              `json:"zlib"`
}

type ApplicationPreferences struct {
    AddTrackers                         string              `json:"add_trackers"`
    AddTrackersEnabled                  bool                `json:"add_trackers_enabled"`
    AltDownloadLimit                    int64               `json:"alt_dl_limit"`
    AltUploadLimit                      int64               `json:"alt_up_limit"`
    AltWebUIEnabled                     bool                `json:"alternative_webui_enabled"`
    AltWebUIPath                        string              `json:"alternative_webui_path"`
    AnnounceIP                          string              `json:"announce_ip"`
    AnnounceToAllTiers                  bool                `json:"announce_to_all_tiers"`
    AnnounceToAllTrackers               bool                `json:"announce_to_all_trackers"`
    AnonymousMode                       bool                `json:"anonymous_mode"`
    AsyncIOThreads                      int64               `json:"async_io_threads"`
    AutoDeleteMode                      int64               `json:"auto_delete_mode"`
    AutoTMMEnabled                      bool                `json:"auto_tmm_enabled"`
    AutoRunEnabled                      bool                `json:"autorun_enabled"`
    AutoRunProgram                      string              `json:"autorun_program"`
    BannedIPs                           string              `json:"banned_IPs"`
    BittorrentProtocol                  int64               `json:"bittorrent_protocol"`
    BypassAuthSubnetWL                  string              `json:"bypass_auth_subnet_whitelist"`
    BypassAuthSubnetWLEnabled           bool                `json:"bypass_auth_subnet_whitelist_enabled"`
    BypassLocalAuth                     bool                `json:"bypass_local_auth"`
    CategoryChangedTMMEnabled           bool                `json:"category_changed_tmm_enabled"`
    CheckingMemoryUse                   int64               `json:"checking_memory_use"`
    CreateSubFolderEnabled              bool                `json:"create_subfolder_enabled"`
    CurrentInterfaceAddress             string              `json:"current_interface_address"`
    CurrentNetworkInterface             string              `json:"current_network_interface"`
    DHT                                 bool                `json:"dht"`
    DiskCache                           int64               `json:"disk_cache"`
    DiskCacheTTL                        int64               `json:"disk_cache_ttl"`
    DownloadLimit                       int64               `json:"dl_limit"`
    DontCountSlowTorrents               bool                `json:"dont_count_slow_torrents"`
    DynDNSDomain                        string              `json:"dyndns_domain"`
    DynDNSEnabled                       bool                `json:"dyndns_enabled"`
    DynDNSPassword                      string              `json:"dyndns_password"`
    DynDNSService                       int64               `json:"dyndns_service"`
    DynDNSUsername                      string              `json:"dyndns_username"`
    EmbeddedTrackerPort                 int64               `json:"embedded_tracker_port"`
    EnableCoalesceReadWrite             bool                `json:"enable_coalesce_read_write"`
    EnabledEmbeddedTracker              bool                `json:"enable_embedded_tracker"`
    EnableMultiConnFromSameIP           bool                `json:"enable_multi_connections_from_same_ip"`
    EnableOSCache                       bool                `json:"enable_os_cache"`
    EnableSuperSeeding                  bool                `json:"enable_super_seeding"`
    EnableUploadSuggestions             bool                `json:"enable_upload_suggestions"`
    Encryption                          int64               `json:"encryption"`
    ExportDirectory                     string              `json:"export_dir"`
    ExportDirectoryFin                  string              `json:"export_dir_fin"`
    FilePoolSize                        int64               `json:"file_pool_size"`
    IncompleteFilesExtension            bool                `json:"incomplete_files_ext"`
    IPFilterEnabled                     bool                `json:"ip_filter_enabled"`
    IPFilterPath                        string              `json:"ip_filter_path"`
    IPFilterTrackers                    bool                `json:"ip_filter_trackers"`
    LimitLanPeers                       bool                `json:"limit_lan_peers"`
    LimitTCPOverhead                    bool                `json:"limit_tcp_overhead"`
    LimitUTPRate                        bool                `json:"limit_utp_rate"`
    ListenPort                          int64               `json:"listen_port"`
    Locale                              string              `json:"locale"`
    LSD                                 bool                `json:"lsd"`
    MailNotificationAuthEnabled         bool                `json:"mail_notification_auth_enabled"`
    MailNotificationEmail               string              `json:"mail_notification_email"`
    MailNotificationEnabled             bool                `json:"mail_notification_enabled"`
    MailNotificationPassword            string              `json:"mail_notification_password"`
    MailNotificationSender              string              `json:"mail_notification_sender"`
    MailNotificationSMTP                string              `json:"mail_notification_smtp"`
    MailNotificationSSLEnabled          bool                `json:"mail_notification_ssl_enabled"`
    MailNotificationUsername            string              `json:"mail_notification_username"`
    MaxActiveDownloads                  int64               `json:"max_active_downloads"`
    MaxActiveTorrents                   int64               `json:"max_active_torrents"`
    MaxActiveUploads                    int64               `json:"max_active_uploads"`
    MaxConnections                      int64               `json:"max_connec"`
    MaxConnectionPerTorrent             int64               `json:"max_connec_per_torrent"`
    MaxRatio                            int64               `json:"max_ratio"`
    MaxRatioACT                         int64               `json:"max_ratio_act"`
    MaxRatioEnabled                     bool                `json:"max_ratio_enabled"`
    MaxSeedingTime                      int64               `json:"max_seeding_time"`
    MaxSeedingTimeEnabled               bool                `json:"max_seeding_time_enabled"`
    MaxUploads                          int64               `json:"max_uploads"`
    MaxUploadsPerTorrent                int64               `json:"max_uploads_per_torrent"`
    OutgoingPortsMax                    int64               `json:"outgoing_ports_max"`
    OutgoingPortsMin                    int64               `json:"outgoing_ports_min"`
    PEX                                 bool                `json:"pex"`
    PreAllocateAll                      bool                `json:"preallocate_all"`
    ProxyAuthEnabled                    bool                `json:"proxy_auth_enabled"`
    ProxyIP                             string              `json:"proxy_ip"`
    ProxyPassword                       string              `json:"proxy_password"`
    ProxyPeerConnection                 bool                `json:"proxy_peer_connections"`
    ProxyPort                           int64               `json:"proxy_port"`
    ProxyTorrentsOnly                   bool                `json:"proxy_torrents_only"`
    ProxyType                           int64               `json:"proxy_type"`
    ProxyUsername                       string              `json:"proxy_username"`
    QueueingEnabled                     bool                `json:"queueing_enabled"`
    RandomPort                          bool                `json:"random_port"`
    ReCheckCompletedTorrents            bool                `json:"recheck_completed_torrents"`
    ResolvePeerCountries                bool                `json:"resolve_peer_countries"`
    RSSAutoDownloadingEnabled           bool                `json:"rss_auto_downloading_enabled"`
    RSSMaxArticlesPerFeed               int64               `json:"rss_max_articles_per_feed"`
    RSSProcessingEnabled                bool                `json:"rss_processing_enabled"`
    RSSRefreshInterval                  int64               `json:"rss_refresh_interval"`
    SavePath                            string              `json:"save_path"`
    SavePathChangedTMMEnabled           bool                `json:"save_path_changed_tmm_enabled"`
    SaveResumeDataInterval              int64               `json:"save_resume_data_interval"`
    ScanDirs                            map[string]bool     `json:"scan_dirs"`
    ScheduleFromHour                    int64               `json:"schedule_from_hour"`
    ScheduleFromMinute                  int64               `json:"schedule_from_min"`
    ScheduleToHour                      int64               `json:"schedule_to_hour"`
    ScheduleToMinute                    int64               `json:"schedule_to_min"`
    SchedulerDays                       int64               `json:"scheduler_days"`
    SchedulerEnabled                    bool                `json:"scheduler_enabled"`
    SendBufferLowWatermark              int64               `json:"send_buffer_low_watermark"`
    SendBufferWatermark                 int64               `json:"send_buffer_watermark"`
    SendBuffetWatermarkFactor           int64               `json:"send_buffer_watermark_factor"`
    SlowTorrentDownloadRateThreshold    int64               `json:"slow_torrent_dl_rate_threshold"`
    SlowTorrentInactiveTimer            int64               `json:"slow_torrent_inactive_timer"`
    SlowTorrentUploadRateThreshold      int64               `json:"slow_torrent_ul_rate_threshold"`
    SocketBacklogSize                   int64               `json:"socket_backlog_size"`
    StartPausedEnabled                  bool                `json:"start_paused_enabled"`
    TemporaryPath                       string              `json:"temp_path"`
    TemporaryPathEnabled                bool                `json:"temp_path_enabled"`
    TorrentChangedTMMEnabled            bool                `json:"torrent_changed_tmm_enabled"`
    UploadLimit                         int64               `json:"up_limit"`
    UploadChokingAlgorithm              int64               `json:"upload_choking_algorithm"`
    UploadSlotsBehavior                 int64               `json:"upload_slots_behavior"`
    UPNP                                bool                `json:"upnp"`
    UseHTTP                             bool                `json:"use_https"`
    UTPTCPMixedMode                     int64               `json:"utp_tcp_mixed_mode"`
    WebUIAddress                        string              `json:"web_ui_address"`
    WebUIClickJackingProtectionEnabled  bool                `json:"web_ui_clickjacking_protection_enabled"`
    WebUICSRFProtectionEnabled          bool                `json:"web_ui_csrf_protection_enabled"`
    WebUIDomainList                     string              `json:"web_ui_domain_list"`
    WebUIHostHeaderValidationEnabled    bool                `json:"web_ui_host_header_validation_enabled"`
    WebUIHttpsCertificatePath           string              `json:"web_ui_https_cert_path"`
    WebUIHttpsKeyPath                   string              `json:"web_ui_https_key_path"`
    WebUIPort                           int64               `json:"web_ui_port"`
    WebUISessionTimeout                 int64               `json:"web_ui_session_timeout"`
    WebUIUPNP                           bool                `json:"web_ui_upnp"`
    WebUIUsername                       string              `json:"web_ui_username"`
}

// Get application version
func (client *Client) ApplicationVersion() (*ApplicationVersion, error) {
    version := &ApplicationVersion{}
    response, err := client.sendGetRequest(ENDPOINT_APPLICATION_VERSION, nil)
    if err != nil {
        return nil, err
    }
    versionString := strings.ReplaceAll(string(client.responseToBytes(response)), "v", "")
    parts := strings.Split(versionString, ".")

    version.Major, err = strconv.ParseUint(parts[0], 10, 32)
    if err != nil {
        return nil, errors.StringToIntError.ToError(err)
    }
    version.Minor, err = strconv.ParseUint(parts[1], 10, 32)
    if err != nil {
        return nil, errors.StringToIntError.ToError(err)
    }
    version.Patch, err = strconv.ParseUint(parts[2], 10, 32)
    if err != nil {
        return nil, errors.StringToIntError.ToError(err)
    }
    return version, nil
}

// Get application api version
func (client *Client) ApplicationAPIVersion() (*ApplicationAPIVersion, error) {
    version := &ApplicationAPIVersion{}
    response, err := client.sendGetRequest(ENDPOINT_APPLICATION_API_VERSION, nil)
    if err != nil {
        return nil, err
    }
    parts := strings.Split(string(client.responseToBytes(response)), ".")

    version.Major, err = strconv.ParseUint(parts[0], 10, 32)
    if err != nil {
        return nil, errors.StringToIntError.ToError(err)
    }
    version.Minor, err = strconv.ParseUint(parts[1], 10, 32)
    if err != nil {
        return nil, errors.StringToIntError.ToError(err)
    }
    return version, nil
}

// Get application build information
func (client *Client) ApplicationBuildInformation() (*ApplicationBuildInfo, error) {
    var buildInformation ApplicationBuildInfo
    result, err := client.get(&buildInformation, ENDPOINT_APPLICATION_BUILD_INFO, nil)
    return result.(*ApplicationBuildInfo), err
}

// Get application preferences
func (client *Client) ApplicationPreferences() (*ApplicationPreferences, error) {
    var applicationPreferences ApplicationPreferences
    result, err := client.get(&applicationPreferences, ENDPOINT_APPLICATION_PREFERENCES, nil)
    return result.(*ApplicationPreferences), err
}

// Get default save path
func (client *Client) ApplicationDefaultSavePath() (string, error) {
    response, err := client.sendGetRequest(ENDPOINT_APPLICATION_DEFAULT_SAVE_PATH, nil)
    if err != nil {
        return "", err
    }
    return string(client.responseToBytes(response)), nil
}
