import { useEffect, useState } from "react";

interface WebSocketHook {
	messages: string[];
	sendMessage: (message: string) => void;
}

const useWebSocket = (url: string): WebSocketHook => {
	const [ws, setWs] = useState<WebSocket | null>(null);
	const [messages, setMessages] = useState<string[]>([]);

	useEffect(() => {
		if (url) {
			try {
				const socket = new WebSocket(url);

				socket.onopen = () => {
					console.log(
						"WebSocket connection opened",
					);
				};

				socket.onmessage = (event) => {
					const data = event.data;
					setMessages((prevMessages) => [
						...prevMessages,
						data,
					]);
					console.log(data);
				};

				socket.onerror = (error) => {
					console.error(
						"WebSocket error:",
						error,
					);
				};

				socket.onclose = () => {
					console.log(
						"WebSocket connection closed",
					);
				};

				setWs(socket);

				return () => {
					socket.close();
				};
			} catch (error) {
				console.error("WebSocket error:", error);
			}
		}
	}, [url]);

	const sendMessage = (message: string) => {
		if (ws && ws.readyState === WebSocket.OPEN) {
			ws.send(message);
		} else {
			console.error("WebSocket is not open");
		}
	};

	return { messages, sendMessage };
};

export default useWebSocket;
