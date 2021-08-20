package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Boolean conversions.
 * 
 * @author Cheejyg
 */
public class Boolean {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Boolean conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * BooleanToString parses <code>boolean</code> boolean_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String BooleanToString(boolean boolean_) {
		if (boolean_ == false) {
			return "false";
		} else if (boolean_ == true) {
			return "true";
		} else {
		}

		return null;
	}

	/**
	 * BooleanToLong parses <code>boolean</code> boolean_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static java.lang.Long BooleanToLong(boolean boolean_) {
		if (boolean_ == false) {
			return 0L;
		} else if (boolean_ == true) {
			return 1L;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToInteger parses <code>boolean</code> boolean_ to
	 * <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static java.lang.Integer BooleanToInteger(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToShort parses <code>boolean</code> boolean_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static java.lang.Short BooleanToShort(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToByte parses <code>boolean</code> boolean_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static java.lang.Byte BooleanToByte(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return null;
	}

	/**
	 * BooleanToBoolean returns <code>Boolean</code> boolean_.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be returned.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean BooleanToBoolean(boolean boolean_) {
		return boolean_;
	}
}
