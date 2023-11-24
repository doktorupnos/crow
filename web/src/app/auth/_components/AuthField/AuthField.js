import styles from "./AuthField.module.css";

export default function AuthField({ id, type, maxlen, placeholder, onchange }) {
	return (
		<div>
			<label
				htmlFor={id}
				className={`
					${styles.authFieldLabel}
					shadow-sm
					focus-within:border-blue-600
					focus-within:ring-1
					focus-within:ring-blue-600 
				`}
			>
				<input
					type={type}
					id={id}
					maxLength={maxlen}
					placeholder={placeholder}
					className={`
						${styles.authFieldInput}
						peer
						placeholder-transparent
						focus:border-transparent
						focus:outline-none
						focus:ring-0
					`}
					onChange={onchange}
					required
				/>
				<span
					className={`
						${styles.authFieldSpan}
						-translate-y-1/2
						transition-all
						peer-placeholder-shown:top-1/2
						peer-placeholder-shown:text-sm
						peer-focus:top-3
						peer-focus:text-xs
					`}
				>
					{placeholder}
				</span>
			</label>
		</div>
	);
}
