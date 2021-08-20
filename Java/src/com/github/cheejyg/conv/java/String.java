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
	 * StringToBoolean parses <code>String</code> string_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean StringToBoolean(java.lang.String string_) {
		// return java.lang.Boolean.parseBoolean(string_);

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
