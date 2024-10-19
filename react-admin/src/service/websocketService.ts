// websocketService.ts

export interface WebSocketDataType {
    Type: string;
}

export interface WebSocketData {
    Version: number;
    ClientId: string;
    Type: string;
    Service: string;
    Api: string;
    Timestamp: number;
    AttemptCount: number;
    Region: string;
    UserAgent: string;
    FinalHttpStatusCode: number;
    Latency: number;
    MaxRetriesExceeded: number;
    FinalAwsException: string;
    FinalAwsExceptionMessage: string;
}
export interface WebSocketAttemptData {
    Version: number,
    ClientId: string,
    Type: string,
    Service: string,
    Api: string,
    Timestamp: number,
    AttemptLatency: number,
    Fqdn: string,
    UserAgent: string,
    AccessKey: string,
    Region: string,
    SessionToken: string
    HttpStatusCode: number,
    XAmznRequestId: string
}

export type WebSocketCounter = {
    counter200: number;
    counter400: number;
}

export type Status = 'Operational' | 'Down' | 'error';

export class WebSocketService {
    private readonly url: string;
    private socket: WebSocket | null;

    private listeners: Array<(data: WebSocketData) => void>;
    private counterListeners: Array<(data: WebSocketCounter) => void>;
    private counterAttemptListeners: Array<(data: WebSocketCounter) => void>;
    private statusListeners: Array<(status: Status) => void>;
    private clientListeners: Array<(count: number) => void>;
    private regionListeners: Array<(count: number) => void>;

    private _counter: WebSocketCounter;
    private _counterAttempt: WebSocketCounter;
    private _clientCounter: Set<string> = new Set();
    private _regionCounter: Set<string> = new Set();

    constructor(url: string) {
        this.url = url;
        this.socket = null;
        this.listeners = [];
        this.statusListeners = [];
        this.counterListeners = [];
        this.counterAttemptListeners = [];
        this.clientListeners = [];
        this.regionListeners = [];

        this._counter = {
            counter200: 0,
            counter400: 0
        }
        this._counterAttempt = {
            counter200: 0,
            counter400: 0
        }
    }

    connect() {
        this.socket = new WebSocket(this.url);

        this.socket.onopen = () => this._notifyStatusListeners('Operational');
        this.socket.onclose = () => this._notifyStatusListeners('Down');
        this.socket.onerror = () => this._notifyStatusListeners('error');
        this.socket.onmessage = (message) => {
            console.log(message.data)
            try {
                const dataType: WebSocketDataType = JSON.parse(message.data);

                if (dataType.Type === 'ApiCall') {
                    const data: WebSocketData = JSON.parse(message.data);

                    // Process status
                    if (data.FinalHttpStatusCode === 200) {
                        this._counter.counter200 = this._counter.counter200 + 1;
                    } else {
                        this._counter.counter400 = this._counter.counter400 + 1;
                    }
                    // Process clients
                    this._clientCounter.add(data.ClientId)
                    this._notifyClientListeners(this._clientCounter.size)

                    // Process region
                    this._regionCounter.add(data.Region)
                    this._notifyRegionListeners(this._regionCounter.size)

                    this._notifyListeners(data);
                    this._notifyCounterListeners(this._counter);
                }

                if (dataType.Type === 'ApiCallAttempt') {

                    const data: WebSocketAttemptData = JSON.parse(message.data);

                    // Check the FinalHttpStatusCode and accumulate counts
                    if (data.HttpStatusCode === 200) {
                        this._counterAttempt.counter200 = this._counterAttempt.counter200 + 1;
                    } else {
                        this._counterAttempt.counter400 = this._counterAttempt.counter400 + 1;
                    }
                    this._notifyListeners(data);
                    this._notifyCounterAttemptListeners(this._counterAttempt);
                }

            } catch (error) {
                console.error('Error parsing WebSocket message:', error);
            }
        };
    }

    private _notifyClientListeners(data: number) {
        this.clientListeners.forEach(listener => listener(data));
    }

    addClientListener(listener: (data: number) => void) {
        this.clientListeners.push(listener);
    }

    removeClientListener(listener: (data: number) => void) {
        this.clientListeners = this.clientListeners.filter(l => l !== listener);
    }

    // region

    private _notifyRegionListeners(data: number) {
        this.regionListeners.forEach(listener => listener(data));
    }

    addRegionListener(listener: (data: number) => void) {
        this.regionListeners.push(listener);
    }

    removeRegionListener(listener: (data: number) => void) {
        this.regionListeners = this.regionListeners.filter(l => l !== listener);
    }


    private _notifyCounterListeners(data: WebSocketCounter) {
        this.counterListeners.forEach(listener => listener(data));
    }

    addCounterListener(listener: (data: WebSocketCounter) => void) {
        this.counterListeners.push(listener);
    }

    removeCounterListener(listener: (data: WebSocketCounter) => void) {
        this.counterListeners = this.counterListeners.filter(l => l !== listener);
    }

    private _notifyCounterAttemptListeners(data: WebSocketCounter) {
        this.counterAttemptListeners.forEach(listener => listener(data));
    }

    addCounterAttemptListener(listener: (data: WebSocketCounter) => void) {
        this.counterAttemptListeners.push(listener);
    }

    removeCounterAttemptListener(listener: (data: WebSocketCounter) => void) {
        this.counterAttemptListeners = this.counterAttemptListeners.filter(l => l !== listener);
    }

    private _notifyListeners(data: any) {
        this.listeners.forEach(listener => listener(data));
    }

    private _notifyStatusListeners(status: Status) {
        this.statusListeners.forEach(listener => listener(status));
    }

    addDataListener(listener: (data: WebSocketData) => void) {
        this.listeners.push(listener);
    }

    addStatusListener(listener: (status: Status) => void) {
        this.statusListeners.push(listener);
    }

    removeDataListener(listener: (data: WebSocketData) => void) {
        this.listeners = this.listeners.filter(l => l !== listener);
    }

    removeStatusListener(listener: (status: Status) => void) {
        this.statusListeners = this.statusListeners.filter(l => l !== listener);
    }
}

const websocketService = new WebSocketService('ws://localhost:8080/ws/apicall');
export default websocketService;
