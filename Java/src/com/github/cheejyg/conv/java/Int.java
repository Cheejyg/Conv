package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Int conversions.
 * 
 * @author Cheejyg
 */
public class Int {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Int conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * IntToString parses <code>int</code> int_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String IntToString(int int_) {
		String string_ = Conv.ToString(int_);

		return string_;
	}

	/**
	 * IntToLong parses <code>int</code> int_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static java.lang.Long IntToLong(int int_) {
		Long long_ = java.lang.Long.valueOf(int_);

		return long_;
	}

	/**
	 * IntToInteger parses <code>int</code> int_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static java.lang.Integer IntToInteger(int int_) {
		return int_;
	}

	/**
	 * IntToShort parses <code>int</code> int_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static java.lang.Short IntToShort(int int_) {
		Short short_ = java.lang.Short.valueOf((short) int_);

		return short_;
	}

	/**
	 * IntToByte parses <code>int</code> int_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static java.lang.Byte IntToByte(int int_) {
		Byte byte_ = java.lang.Byte.valueOf((byte) int_);

		return byte_;
	}

	/**
	 * IntToBoolean parses <code>int</code> int_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean IntToBoolean(int int_) {
		if (int_ == 0) {
			return false;
		} else if (int_ == 1) {
			return true;
		} else {
		}

		return null;
	}
}
