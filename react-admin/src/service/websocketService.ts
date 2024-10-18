// websocketService.ts
export interface WebSocketData {
    id: number;
    name: string;
    status: string;
    createdAt: string;
}

export type Status = 'Operational' | 'Down' | 'error';

export class WebSocketService {
    private url: string;
    private socket: WebSocket | null;
    private listeners: Array<(data: WebSocketData) => void>;
    private statusListeners: Array<(status: Status) => void>;

    constructor(url: string) {
        this.url = url;
        this.socket = null;
        this.listeners = [];
        this.statusListeners = [];
    }

    connect() {
        this.socket = new WebSocket(this.url);

        this.socket.onopen = () => this._notifyStatusListeners('Operational');
        this.socket.onclose = () => this._notifyStatusListeners('Down');
        this.socket.onerror = () => this._notifyStatusListeners('error');
        this.socket.onmessage = (message) => {
            const data: WebSocketData = JSON.parse(message.data);
            this._notifyListeners(data);
        };
    }

    private _notifyListeners(data: WebSocketData) {
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

const websocketService = new WebSocketService('ws://localhost:8080/ws');
export default websocketService;
