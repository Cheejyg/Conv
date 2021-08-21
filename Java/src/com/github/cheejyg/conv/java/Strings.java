package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * String conversions.
 * 
 * @author Cheejyg
 */
public class Strings {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Strings conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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
			java.lang.Long long_ = String.StringToLong(strings_[i]);
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
			java.lang.Integer integer_ = String.StringToInteger(strings_[i]);
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
			java.lang.Short short_ = String.StringToShort(strings_[i]);
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
			java.lang.Byte byte_ = String.StringToByte(strings_[i]);
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
			java.lang.Boolean boolean_ = String.StringToBoolean(strings_[i]);
			if (boolean_ == null) {
				return null;
			}

			booleans_[i] = boolean_;
		}

		return booleans_;
	}
}
