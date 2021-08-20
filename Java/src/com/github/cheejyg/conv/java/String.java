package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * String conversions.
 * 
 * @author Cheejyg
 */
public class String {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// String conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * StringToLong parses <code>String</code> string_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static java.lang.Long StringToLong(java.lang.String string_) {
		java.lang.Long long_;

		try {
			long_ = java.lang.Long.valueOf(string_);
		} catch (NumberFormatException e) {
			return null;
		}

		return long_;
	}

	/**
	 * StringToInteger parses <code>String</code> string_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static java.lang.Integer StringToInteger(java.lang.String string_) {
		java.lang.Integer integer_;

		try {
			integer_ = java.lang.Integer.valueOf(string_);
		} catch (NumberFormatException e) {
			return null;
		}

		return integer_;
	}

	/**
	 * StringToShort parses <code>String</code> string_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static java.lang.Short StringToShort(java.lang.String string_) {
		java.lang.Short short_;

		try {
			short_ = java.lang.Short.valueOf(string_);
		} catch (NumberFormatException e) {
			return null;
		}

		return short_;
	}

	/**
	 * StringToByte parses <code>String</code> string_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static java.lang.Byte StringToByte(java.lang.String string_) {
		java.lang.Byte byte_;

		try {
			byte_ = java.lang.Byte.valueOf(string_);
		} catch (NumberFormatException e) {
			return null;
		}

		return byte_;
	}

	/**
	 * StringToBoolean parses <code>String</code> string_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean StringToBoolean(java.lang.String string_) {
		// return java.lang.Boolean.valueOf(string_);

		switch (string_.trim()) {
		case "0":
		case "f":
		case "F":
		case "FALSE":
		case "false":
		case "False":
			return false;
		case "1":
		case "t":
		case "T":
		case "TRUE":
		case "true":
		case "True":
			return true;
		default:
		}

		return null;
	}
}
