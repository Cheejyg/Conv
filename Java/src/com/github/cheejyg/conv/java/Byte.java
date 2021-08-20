package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Byte conversions.
 * 
 * @author Cheejyg
 */
public class Byte {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Byte conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ByteToString parses <code>byte</code> byte_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String ByteToString(byte byte_) {
		java.lang.String string_ = Conv.ToString(byte_);

		return string_;
	}

	/**
	 * ByteToLong parses <code>byte</code> byte_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static java.lang.Long ByteToLong(byte byte_) {
		java.lang.Long long_ = java.lang.Long.valueOf(byte_);

		return long_;
	}

	/**
	 * ByteToInteger parses <code>byte</code> byte_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static java.lang.Integer ByteToInteger(byte byte_) {
		java.lang.Integer integer_ = java.lang.Integer.valueOf(byte_);

		return integer_;
	}

	/**
	 * ByteToShort parses <code>byte</code> byte_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static java.lang.Short ByteToShort(byte byte_) {
		java.lang.Short short_ = java.lang.Short.valueOf(byte_);

		return short_;
	}

	/**
	 * ByteToByte returns <code>Byte</code> byte_.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static java.lang.Byte ByteToByte(byte byte_) {
		return byte_;
	}

	/**
	 * ByteToBoolean parses <code>byte</code> byte_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean ByteToBoolean(byte byte_) {
		if (byte_ == 0) {
			return false;
		} else if (byte_ == 1) {
			return true;
		} else {
		}

		return null;
	}
}
