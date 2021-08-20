package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Short conversions.
 * 
 * @author Cheejyg
 */
public class Short {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Short conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ShortToString parses <code>short</code> short_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String ShortToString(short short_) {
		java.lang.String string_ = Conv.ToString(short_);

		return string_;
	}

	/**
	 * ShortToLong parses <code>short</code> short_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static java.lang.Long ShortToLong(short short_) {
		java.lang.Long long_ = java.lang.Long.valueOf(short_);

		return long_;
	}

	/**
	 * ShortToInteger parses <code>short</code> short_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static java.lang.Integer ShortToInteger(short short_) {
		java.lang.Integer integer_ = java.lang.Integer.valueOf(short_);

		return integer_;
	}

	/**
	 * ShortToShort parses <code>short</code> short_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static java.lang.Short ShortToShort(short short_) {
		return short_;
	}

	/**
	 * ShortToByte parses <code>short</code> short_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static java.lang.Byte ShortToByte(short short_) {
		java.lang.Byte byte_ = java.lang.Byte.valueOf((byte) short_);

		return byte_;
	}

	/**
	 * ShortToBoolean parses <code>short</code> short_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean ShortToBoolean(short short_) {
		if (short_ == 0) {
			return false;
		} else if (short_ == 1) {
			return true;
		} else {
		}

		return null;
	}
}
