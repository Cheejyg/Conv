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
	 * StringToString returns <code>String</code> string_.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be returned.
	 * @return <code>String</code>
	 */
	public static java.lang.String StringToString(java.lang.String string_) {
		return string_;
	}

	/**
	 * StringToStringArray parses <code>String</code> string_ to
	 * <code>String[]</code> with provided regex.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @param regex_  The provided regex separator.
	 * @return <code>String[]</code>
	 */
	public static java.lang.String[] StringToStringArray(java.lang.String string_, java.lang.String regex_) {
		java.lang.String[] strings_ = string_.split(regex_);

		return strings_;
	}

	/**
	 * StringToLongArray parses <code>String</code> string_ to <code>Long[]</code>
	 * with provided regex.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @param regex_  The provided regex separator.
	 * @return <code>Long[]</code>
	 */
	public static java.lang.Long[] StringToLongArray(java.lang.String string_, java.lang.String regex_) {
		java.lang.String[] strings_ = string_.split(regex_);

		java.lang.Long[] longs_ = new java.lang.Long[strings_.length];

		for (int i = 0; i < strings_.length; i++) {
			java.lang.Long long_ = StringToLong(strings_[i]);
			if (long_ == null) {
				return null;
			}

			longs_[i] = long_;
		}

		return longs_;
	}

	/**
	 * StringToIntegerArray parses <code>String</code> string_ to
	 * <code>Integer[]</code> with provided regex.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @param regex_  The provided regex separator.
	 * @return <code>Integer[]</code>
	 */
	public static java.lang.Integer[] StringToIntegerArray(java.lang.String string_, java.lang.String regex_) {
		java.lang.String[] strings_ = string_.split(regex_);

		java.lang.Integer[] integers_ = new java.lang.Integer[strings_.length];

		for (int i = 0; i < strings_.length; i++) {
			java.lang.Integer integer_ = StringToInteger(strings_[i]);
			if (integer_ == null) {
				return null;
			}

			integers_[i] = integer_;
		}

		return integers_;
	}

	/**
	 * StringToShortArray parses <code>String</code> string_ to <code>Short[]</code>
	 * with provided regex.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @param regex_  The provided regex separator.
	 * @return <code>Short[]</code>
	 */
	public static java.lang.Short[] StringToShortArray(java.lang.String string_, java.lang.String regex_) {
		java.lang.String[] strings_ = string_.split(regex_);

		java.lang.Short[] shorts_ = new java.lang.Short[strings_.length];

		for (int i = 0; i < strings_.length; i++) {
			java.lang.Short short_ = StringToShort(strings_[i]);
			if (short_ == null) {
				return null;
			}

			shorts_[i] = short_;
		}

		return shorts_;
	}

	/**
	 * StringToByteArray parses <code>String</code> string_ to <code>Byte[]</code>
	 * with provided regex.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @param regex_  The provided regex separator.
	 * @return <code>Byte[]</code>
	 */
	public static java.lang.Byte[] StringToByteArray(java.lang.String string_, java.lang.String regex_) {
		java.lang.String[] strings_ = string_.split(regex_);

		java.lang.Byte[] bytes_ = new java.lang.Byte[strings_.length];

		for (int i = 0; i < strings_.length; i++) {
			java.lang.Byte byte_ = StringToByte(strings_[i]);
			if (byte_ == null) {
				return null;
			}

			bytes_[i] = byte_;
		}

		return bytes_;
	}

	/**
	 * StringToBooleanArray parses <code>String</code> string_ to
	 * <code>Boolean[]</code> with provided regex.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @param regex_  The provided regex separator.
	 * @return <code>Boolean[]</code>
	 */
	public static java.lang.Boolean[] StringToBooleanArray(java.lang.String string_, java.lang.String regex_) {
		java.lang.String[] strings_ = string_.split(regex_);

		java.lang.Boolean[] booleans_ = new java.lang.Boolean[strings_.length];

		for (int i = 0; i < strings_.length; i++) {
			java.lang.Boolean boolean_ = StringToBoolean(strings_[i]);
			if (boolean_ == null) {
				return null;
			}

			booleans_[i] = boolean_;
		}

		return booleans_;
	}

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
