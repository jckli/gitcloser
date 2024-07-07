import { useRef, useCallback } from "react";

function useDebounce(callback: (...args: any[]) => void, delay: number) {
	const timeoutRef = useRef<number | NodeJS.Timeout | null>(null);

	const debouncedCallback = useCallback(
		(...args: any[]) => {
			if (timeoutRef.current) {
				clearTimeout(
					timeoutRef.current as NodeJS.Timeout,
				);
			}
			timeoutRef.current = setTimeout(() => {
				callback(...args);
			}, delay);
		},
		[callback, delay],
	);

	return debouncedCallback;
}

export default useDebounce;
